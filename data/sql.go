package data

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const DBTYPE string = "sqlite3"
const DBPATH string = "./db/database.db"

func CreateTable() {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	// Create the activities table if it does not exist.
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS activities (id INTEGER PRIMARY KEY, name TEXT UNIQUE)")
	statement.Exec()

	// Create the records table if it does not exist.
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, activity TEXT, date DATE, timepassed TIME)")
	statement.Exec()
}

func GetActivitiesString() [][]string {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM activities")

	if err != nil {
		log.Fatal("There was a problem with the query.")
	}

	var activities [][]string
	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		activities = append(activities, []string{strconv.Itoa(id), name})
	}

	return activities
}

func GetRecordsString() [][]string {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM records")

	if err != nil {
		log.Fatal("There was a problem with the query.")
	}

	var records [][]string
	for rows.Next() {
		var id int
		var name string
		var datet string
		var timepassed int

		if err := rows.Scan(&id, &name, &datet, &timepassed); err != nil {
			log.Fatal(err)
		}

		records = append(records, []string{strconv.Itoa(id), name, datet, strconv.Itoa(timepassed)})
	}

	return records
}

func AddActivity(name string) {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO activities (name) VALUES (?)")
	statement.Exec(name)
}

func AddRecord(name string, date string, timepassed int) {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO records (activity, date, timepassed) VALUES (?, ?, ?)")
	statement.Exec(name, date, timepassed)
}

func DeleteActivity(name string) {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	statement, _ := db.Prepare("DELETE FROM activities WHERE name = ?;")
	statement.Exec(name)

	// Deletes the records pretaining to the activity deleted
	statement, _ = db.Prepare("DELETE FROM records WHERE activity = ?")
	statement.Exec(name)
}

func DeleteRecord() {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	statement, _ := db.Prepare("DELETE FROM records WHERE activity = ?")
	statement.Exec()
}
