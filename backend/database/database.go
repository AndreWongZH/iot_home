package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const databaseFilePath = "iothome.db"

const createRooms string = `
	CREATE TABLE IF NOT EXISTS rooms (
		room_id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
	)
`

const createDevInfo string = `
	CREATE TABLE IF NOT EXISTS deviceInfo (
		room_id INTEGER NOT NULL,
		device_id INTEGER NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		ipaddr TEXT NOT NULL,
		type TEXT NOT NULL,
		hostname TEXT,
		macaddress TEXT,
		FOREIGN KEY (room_id)
			REFERENCES rooms (room_id)
	)
`

const createDevStatus string = `
	CREATE TABLE IF NOT EXISTS deviceStatus (
		device_status_id INTEGER NOT NULL PRIMARY KEY,
		device_id INTEGER NOT NULL,
		status INTEGER NOT NULL,
		on INTEGER NOT NULL,
		FOREIGN KEY (device_id)
			REFERENCES rooms (device_id)
	)
`

func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", databaseFilePath)

	if err != nil {
		log.Println("Failed to open sqlite3 database")
	}

	if _, err := db.Exec(createRooms); err != nil {
		log.Println("Failed to create new rooms table")
	}

	if _, err := db.Exec(createDevInfo); err != nil {
		log.Println("Failed to create new devinfo table")
	}

	if _, err := db.Exec(createDevStatus); err != nil {
		log.Println("Failed to create new devstatus table")
	}

	return db
}
