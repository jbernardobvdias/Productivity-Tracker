package data

import (
	"database/sql"
	"log"
	"prod_tracker/model"

	_ "github.com/mattn/go-sqlite3"
)

const DBTYPE string = "sqlite3"
const DBPATH string = "./db/database.db"

func CreateTable() {
	db, err := sql.Open(DBTYPE, DBPATH)

	if err != nil {
		log.Fatal("There was an error connecting to the database.", err.Error())
	}

	defer db.Close()

	// Create the activities table if it does not exist.
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS activities (id INTEGER PRIMARY KEY, name TEXT)")
	statement.Exec()

	// Create the records table if it does not exist.
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, activity TEXT, date DATE, timepassed TIME)")
	statement.Exec()
}

func LoadTable() ([]model.Activity, []model.Record) {
	db, err := sql.Open(DBTYPE, DBPATH)

	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}

	defer db.Close()

	// Load the activities table

	rows, err := db.Query("SELECT * FROM activities")

	if err != nil {
		log.Fatal("There was a problem with the querry.")
	}

	var activities []model.Activity
	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		activities = append(activities, model.Activity{Id: id, Name: name})
	}

	// Load the records table

	rows, err = db.Query("SELECT * FROM records")

	if err != nil {
		log.Fatal("There was a problem with the querry.")
	}

	var records []model.Record
	for rows.Next() {

		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		records = append(records, model.Record{Id: id, ActivityName: name, DateT: name, TimePassed: id})
	}

	return activities, records
}

func AddActivity() {
	db, err := sql.Open(DBTYPE, DBPATH)

	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}

	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO activities (name) VALUES (?)")
	statement.Exec()
}

func AddRecord() {
	db, err := sql.Open(DBTYPE, DBPATH)

	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}

	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO records (activity, date, timepassed) VALUES (?, ?, ?)")
	statement.Exec()
}
