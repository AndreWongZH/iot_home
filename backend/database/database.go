package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/AndreWongZH/iothome/models"
	_ "github.com/mattn/go-sqlite3"
)

const databaseFilePath = "iothome.db"

const createRooms string = `
	CREATE TABLE IF NOT EXISTS rooms (
		room_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT
	)
`

const createDevInfo string = `
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

const createDevStatus string = `
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

const getRoomId string = `SELECT room_id FROM rooms WHERE name=?`
const insertNewRoom string = `INSERT INTO rooms (name) VALUES (?)`
const getRooms string = `SELECT rooms.name, count(deviceInfo.room_id) FROM rooms LEFT JOIN deviceInfo ON rooms.room_id = deviceInfo.room_id GROUP BY rooms.name`

const ifRoomExist string = `SELECT COUNT(*) FROM rooms WHERE name=?`
const ifIpExist string = `SELECT COUNT(*) FROM deviceInfo WHERE ipaddr=?`
const ifIpExistInRoom string = `SELECT COUNT(*) FROM deviceInfo WHERE ipaddr=? AND room_id=?`

const insertNewDevInfo string = `INSERT INTO deviceInfo (room_id, name, ipaddr, type) VALUES (?, ? ,? ,?)`
const insertNewDevStatus string = `INSERT INTO deviceStatus (device_id, connected, on_state) VALUES (?, ? ,?)`

const getDeviceInfoByRoom string = `SELECT dev.name, dev.ipaddr, dev.type, dev.connected, dev.on_state FROM rooms JOIN (SELECT * FROM deviceInfo JOIN deviceStatus WHERE deviceInfo.device_id = deviceStatus.device_id) as dev WHERE rooms.room_id = dev.room_id and rooms.name=?`
const getDeviceInfo string = `SELECT dev.device_id, dev.name, dev.type, dev.connected, dev.on_state FROM rooms JOIN (SELECT * FROM deviceInfo JOIN deviceStatus WHERE deviceInfo.device_id = deviceStatus.device_id) as dev WHERE rooms.room_id = dev.room_id and rooms.name=? and dev.ipaddr=?`

const updateDeviceStatus string = `UPDATE deviceStatus SET on_state=? WHERE device_id=?`

const deleteNewRoom string = `DELETE FROM rooms WHERE name='?'`

func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", databaseFilePath)
	if err != nil {
		log.Println("Failed to open sqlite3 database")
	}

	if _, err := db.Exec("PRAGMA foreign_keys = 1"); err != nil {
		log.Println(err)
	}

	if _, err := db.Exec(createRooms); err != nil {
		log.Println("Failed to create new rooms table")
		log.Println(err)
	}

	if _, err := db.Exec(createDevInfo); err != nil {
		log.Println("Failed to create new devinfo table")
		log.Println(err)
	}

	if _, err := db.Exec(createDevStatus); err != nil {
		log.Println("Failed to create new devstatus table")
		log.Println(err)
	}

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

func (s *DatabaseManager) AddRoom(room models.RoomInfo) error {
	_, err := s.Db.Exec(insertNewRoom, room.Name)

	if err != nil {
		log.Println("error inserting entry into table")
		log.Println(err)
		return err
	}

	return nil
}

func (s *DatabaseManager) DelRoom(roomname string) {
	_, err := s.Db.Exec(deleteNewRoom, roomname)

	if err != nil {
		log.Println("error deleting entry into table")
		log.Println(err)
	}
}

func (s *DatabaseManager) AddDevice(dev models.RegisteredDevice, devStatus models.DeviceStatus, roomName string) error {
	row := s.Db.QueryRow(getRoomId, roomName)
	var roomId int
	if err := row.Scan(&roomId); err == sql.ErrNoRows {
		log.Println("room_id not found")
		return err
	}

	res, err := s.Db.Exec(insertNewDevInfo, roomId, dev.Name, dev.Ipaddr, dev.Type)

	if err != nil {
		log.Println("error inserting entry into deviceInfo table")
		log.Println(err)
		return err
	}

	var devId int64
	devId, _ = res.LastInsertId()

	_, err = s.Db.Exec(insertNewDevStatus, devId, devStatus.Connected, devStatus.On_state)

	if err != nil {
		log.Println("error inserting entry into deviceStatus table")
		log.Println(err)
		return err
	}

	return nil
}

func (s *DatabaseManager) UpdateDevStatus(device_id int, on_state int) error {
	_, err := s.Db.Exec(updateDeviceStatus, on_state, device_id)
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
	if err := row.Scan(&device_id, &devInfo.Name, &devInfo.Type, &devStatus.Connected, &devStatus.On_state); err == sql.ErrNoRows {
		return device_id, devInfo, devStatus, err
	}

	return device_id, devInfo, devStatus, nil
}

func (s *DatabaseManager) GetDevices(roomName string) ([]models.RegisteredDevice, map[string]models.DeviceStatus, error) {
	rows, err := s.Db.Query(getDeviceInfoByRoom, roomName)
	if err != nil {
		log.Println("error getting entrys from table")
		log.Println(err)
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
			log.Println("failed to scan db rows")
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
		log.Println("error getting entrys from table")
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	rooms := []models.RoomInfo{}
	for rows.Next() {

		ri := models.RoomInfo{}
		err = rows.Scan(&ri.Name, &ri.Count)
		if err != nil {
			log.Println("failed to scan db rows")
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
		return false, err
	}

	if nameCount > 0 {
		return true, nil
	}

	return false, nil
}

func (s *DatabaseManager) CheckIpExist(ipaddr string) (bool, error) {
	row := s.Db.QueryRow(ifIpExist, ipaddr)

	var ipCount int

	if err := row.Scan(&ipCount); err == sql.ErrNoRows {
		return false, err
	}

	fmt.Println(ipCount)

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
		return false, err
	}

	fmt.Println(ipCount)

	if ipCount > 0 {
		return true, nil
	}

	return false, nil
}

func PopulateDatabase() {

	Dbman.AddRoom(models.RoomInfo{Name: "kekw"})
	Dbman.AddRoom(models.RoomInfo{Name: "bedroom"})
	// Dbman.AddDevice(models.RegisteredDevice{Name: "andre", Type: "wled", Ipaddr: "192.168.1.1", Hostname: "123"}, "kekw")
	// Dbman.AddDevice(models.RegisteredDevice{Name: "betty", Type: "wled", Ipaddr: "192.168.1.2", Hostname: "456"}, "kekw")
	// Dbman.AddDevice(models.RegisteredDevice{Name: "cathy", Type: "wled", Ipaddr: "192.168.1.3", Hostname: "789"}, "kekw")
	// Dbman.AddDevice(models.RegisteredDevice{Name: "deuick", Type: "wled", Ipaddr: "192.168.1.4", Hostname: "023"}, "kekw")
	// Dbman.AddDevice(models.RegisteredDevice{Name: "andre", Type: "wled", Ipaddr: "192.168.1.1", Hostname: "123"}, "bedroom")
	// Dbman.AddDevice(models.RegisteredDevice{Name: "betty", Type: "wled", Ipaddr: "192.168.1.2", Hostname: "456"}, "bedroom")
	// Dbman.AddDevice(models.RegisteredDevice{Name: "cathy", Type: "wled", Ipaddr: "192.168.1.3", Hostname: "789"}, "bedroom")
	// Dbman.AddDevice(models.RegisteredDevice{Name: "deuick", Type: "wled", Ipaddr: "192.168.1.4", Hostname: "023"}, "bedroom")
}
