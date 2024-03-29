package database

import (
	"database/sql"
	"errors"

	"github.com/AndreWongZH/iothome/logger"
	"github.com/AndreWongZH/iothome/models"
	_ "github.com/mattn/go-sqlite3"
)

const databaseFilePath = "iothome.db"

const (
	createRooms string = `
		CREATE TABLE IF NOT EXISTS rooms (
			room_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE
		)
	`

	createDevInfo string = `
		CREATE TABLE IF NOT EXISTS deviceInfo (
			room_id INTEGER NOT NULL,
			device_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			ipaddr TEXT NOT NULL,
			type TEXT NOT NULL,
			hostname TEXT,
			macaddress TEXT,
			FOREIGN KEY (room_id)
				REFERENCES rooms (room_id)
					ON DELETE CASCADE
		)
	`

	createDevStatus string = `
		CREATE TABLE IF NOT EXISTS deviceStatus (
			device_status_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			device_id INTEGER NOT NULL,
			connected INTEGER NOT NULL,
			on_state INTEGER NOT NULL,
			FOREIGN KEY (device_id)
				REFERENCES deviceInfo (device_id)
					ON DELETE CASCADE
		)
	`

	createUsers string = `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT NOT NULL UNIQUE,
		hash TEXT NOT NULL
		)
	`
)

const (
	insertUser     string = `INSERT INTO users (username, hash) VALUES (?,?)`
	checkUserExist string = `SELECT hash FROM users WHERE username=?`

	getRoomId     string = `SELECT room_id FROM rooms WHERE name=?`
	insertNewRoom string = `INSERT INTO rooms (name) VALUES (?)`
	getRooms      string = `SELECT rooms.name, count(deviceInfo.room_id) FROM rooms LEFT JOIN deviceInfo ON rooms.room_id = deviceInfo.room_id GROUP BY rooms.name`

	ifRoomExist     string = `SELECT COUNT(*) FROM rooms WHERE name=?`
	ifIpExist       string = `SELECT COUNT(*) FROM deviceInfo WHERE ipaddr=?`
	ifIpExistInRoom string = `SELECT COUNT(*) FROM deviceInfo WHERE ipaddr=? AND room_id=?`

	insertNewDevInfo   string = `INSERT INTO deviceInfo (room_id, name, ipaddr, type) VALUES (?, ? ,? ,?)`
	insertNewDevStatus string = `INSERT INTO deviceStatus (device_id, connected, on_state) VALUES (?, ? ,?)`

	getDeviceInfoByRoom string = `SELECT dev.name, dev.ipaddr, dev.type, dev.connected, dev.on_state FROM rooms JOIN (SELECT * FROM deviceInfo JOIN deviceStatus WHERE deviceInfo.device_id = deviceStatus.device_id) as dev WHERE rooms.room_id = dev.room_id and rooms.name=?`
	getDeviceInfo       string = `SELECT dev.device_id, dev.name, dev.type, dev.connected, dev.on_state FROM rooms JOIN (SELECT * FROM deviceInfo JOIN deviceStatus WHERE deviceInfo.device_id = deviceStatus.device_id) as dev WHERE rooms.room_id = dev.room_id and rooms.name=? and dev.ipaddr=?`

	updateDeviceStatus string = `UPDATE deviceStatus SET connected=?, on_state=? WHERE device_id=?`

	deleteRoom   string = `DELETE FROM rooms WHERE name=?`
	deleteDevice string = `DELETE FROM deviceInfo WHERE ipaddr IN (
		SELECT ipaddr FROM rooms JOIN deviceInfo ON rooms.room_id=deviceInfo.room_id WHERE rooms.name=? and deviceInfo.ipaddr=?
		)
	`
)

func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", databaseFilePath)
	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database")
	}

	if _, err := db.Exec("PRAGMA foreign_keys = 1"); err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database")
	}

	if _, err := db.Exec(createRooms); err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "failed to create new rooms table")
	}

	if _, err := db.Exec(createDevInfo); err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "failed to create new devinfo table")
	}

	if _, err := db.Exec(createDevStatus); err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "failed to create new devstatus table")
	}

	if _, err := db.Exec(createUsers); err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "failed to create new users table")
	}

	logger.SugarLog.Info("database initialized")

	return db
}

type DatabaseManager struct {
	Db *sql.DB
}

var Dbman *DatabaseManager

func InitializeGlobals(db *sql.DB) {
	Dbman = &DatabaseManager{
		Db: db,
	}
}

func (s *DatabaseManager) AddUser(username string, hash string) error {
	row := s.Db.QueryRow(checkUserExist, username)
	var hashdb string
	if err := row.Scan(&hashdb); err == sql.ErrNoRows {

		_, err := s.Db.Exec(insertUser, username, hash)
		if err != nil {
			logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "add user query failed")
			return err
		}

	} else {
		return errors.New("username already in use")
	}

	return nil
}

func (s *DatabaseManager) QueryUserHash(username string) ([]byte, error) {
	row := s.Db.QueryRow(checkUserExist, username)
	var hashdb string
	if err := row.Scan(&hashdb); err == sql.ErrNoRows {
		return nil, errors.New("username not found")
	}

	return []byte(hashdb), nil
}

func (s *DatabaseManager) AddRoom(room models.RoomInfo) error {
	_, err := s.Db.Exec(insertNewRoom, room.Name)

	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "insertNewRoom query failed")
		return err
	}

	return nil
}

func (s *DatabaseManager) DelRoom(roomname string) error {
	_, err := s.Db.Exec(deleteRoom, roomname)

	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "deleteRoom query failed")
		return err
	}

	return nil
}

func (s *DatabaseManager) DelDevice(roomname string, ipAddr string) error {
	_, err := s.Db.Exec(deleteDevice, roomname, ipAddr)

	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "deleteDevice query failed")
		return err
	}

	return nil
}

func (s *DatabaseManager) AddDevice(dev models.RegisteredDevice, devStatus models.DeviceStatus, roomName string) error {
	row := s.Db.QueryRow(getRoomId, roomName)
	var roomId int
	err := row.Scan(&roomId)
	if err == sql.ErrNoRows {
		return err
	} else if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "getRoomId query failed")
		return err
	}

	res, err := s.Db.Exec(insertNewDevInfo, roomId, dev.Name, dev.Ipaddr, dev.Type)

	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "inserting entry into deviceInfo table")
		return err
	}

	var devId int64
	devId, _ = res.LastInsertId()

	_, err = s.Db.Exec(insertNewDevStatus, devId, devStatus.Connected, devStatus.On_state)

	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "inserting entry into deviceStatus table")
		return err
	}

	return nil
}

func (s *DatabaseManager) UpdateDevStatus(roomName string, ipAddr string, devStatus models.DeviceStatus) error {
	device_id, _, _, err := Dbman.GetDevice(roomName, ipAddr)
	if err != nil {
		return err
	}

	_, err = s.Db.Exec(updateDeviceStatus, devStatus.Connected, devStatus.On_state, device_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseManager) GetDevice(roomName string, ipAddr string) (int, models.RegisteredDevice, models.DeviceStatus, error) {
	row := s.Db.QueryRow(getDeviceInfo, roomName, ipAddr)

	devInfo := models.RegisteredDevice{}
	devStatus := models.DeviceStatus{}
	var device_id int
	err := row.Scan(&device_id, &devInfo.Name, &devInfo.Type, &devStatus.Connected, &devStatus.On_state)
	if err == sql.ErrNoRows {
		return device_id, devInfo, devStatus, err
	} else if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "failed to scan db rows")
		return device_id, devInfo, devStatus, err
	}

	devInfo.Ipaddr = ipAddr

	return device_id, devInfo, devStatus, nil
}

func (s *DatabaseManager) GetDevices(roomName string) ([]models.RegisteredDevice, map[string]models.DeviceStatus, error) {
	rows, err := s.Db.Query(getDeviceInfoByRoom, roomName)
	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "getDeviceInfoByRoom query failed")
		return nil, nil, err
	}
	defer rows.Close()

	devList := []models.RegisteredDevice{}
	devStatus := make(map[string]models.DeviceStatus)
	for rows.Next() {
		rd := models.RegisteredDevice{}
		ds := models.DeviceStatus{}
		err = rows.Scan(&rd.Name, &rd.Ipaddr, &rd.Type, &ds.Connected, &ds.On_state)
		if err != nil {
			logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "failed to scan db rows")
			return nil, nil, err
		}

		devList = append(devList, rd)
		devStatus[rd.Ipaddr] = ds
	}

	return devList, devStatus, nil
}

func (s *DatabaseManager) GetRooms() ([]models.RoomInfo, error) {
	rows, err := s.Db.Query(getRooms)
	if err != nil {
		logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "getRooms query failed")
		return nil, err
	}
	defer rows.Close()

	rooms := []models.RoomInfo{}
	for rows.Next() {

		ri := models.RoomInfo{}
		err = rows.Scan(&ri.Name, &ri.Count)
		if err != nil {
			logger.SugarLog.Errorw(err.Error(), "location", "database", "extra", "failed to scan db rows")
			return nil, err
		}

		rooms = append(rooms, ri)
	}

	return rooms, nil
}

func (s *DatabaseManager) CheckRoomExist(roomName string) (bool, error) {
	row := s.Db.QueryRow(ifRoomExist, roomName)

	var nameCount int

	if err := row.Scan(&nameCount); err == sql.ErrNoRows {
		// room does not exist
		return false, err
	}

	// room exists
	if nameCount > 0 {
		return true, nil
	}

	return false, nil
}

func (s *DatabaseManager) CheckIpExist(ipaddr string) (bool, error) {
	row := s.Db.QueryRow(ifIpExist, ipaddr)

	var ipCount int

	if err := row.Scan(&ipCount); err == sql.ErrNoRows {
		// ip does not exist
		return false, err
	}

	// ip exists
	if ipCount > 0 {
		return true, nil
	}

	return false, nil
}

func (s *DatabaseManager) CheckIpExistInRoom(ipaddr string, roomName string) (bool, error) {
	row := s.Db.QueryRow(getRoomId, roomName)
	var roomId int
	if err := row.Scan(&roomId); err == sql.ErrNoRows {
		return false, errors.New("room id not found")
	}

	row = s.Db.QueryRow(ifIpExistInRoom, ipaddr, roomId)

	var ipCount int

	if err := row.Scan(&ipCount); err == sql.ErrNoRows {
		// ip does not exist
		return false, err
	}

	// ip exists
	if ipCount > 0 {
		return true, nil
	}

	return false, nil
}

// for testing purposes
func PopulateDatabase() {

	Dbman.AddRoom(models.RoomInfo{Name: "kekw"})
	Dbman.AddRoom(models.RoomInfo{Name: "bedroom"})
	Dbman.AddDevice(
		models.RegisteredDevice{Name: "andre", Type: "wled", Ipaddr: "192.168.1.1", Hostname: "123"},
		models.DeviceStatus{Connected: false, On_state: false},
		"kekw",
	)
	Dbman.AddDevice(
		models.RegisteredDevice{Name: "betty", Type: "switch", Ipaddr: "192.168.1.2", Hostname: "123"},
		models.DeviceStatus{Connected: false, On_state: false},
		"kekw",
	)
	Dbman.AddDevice(
		models.RegisteredDevice{Name: "cathy", Type: "wled", Ipaddr: "192.168.1.3", Hostname: "123"},
		models.DeviceStatus{Connected: false, On_state: false},
		"kekw",
	)
	Dbman.AddDevice(
		models.RegisteredDevice{Name: "derick", Type: "switch", Ipaddr: "192.168.1.4", Hostname: "123"},
		models.DeviceStatus{Connected: false, On_state: false},
		"kekw",
	)
	Dbman.AddDevice(
		models.RegisteredDevice{Name: "eagle", Type: "wled", Ipaddr: "192.168.1.5", Hostname: "123"},
		models.DeviceStatus{Connected: false, On_state: false},
		"kekw",
	)
}
