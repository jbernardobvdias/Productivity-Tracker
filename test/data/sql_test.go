package data_test

import (
	"database/sql"
	"prod_tracker/data"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func createTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Create test tables
	_, err = db.Exec(`CREATE TABLE activities (id INTEGER PRIMARY KEY, name TEXT UNIQUE);`)
	if err != nil {
		t.Fatalf("Failed to create activities table: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE records (id INTEGER PRIMARY KEY, activity TEXT, date DATE, timepassed TIME);`)
	if err != nil {
		t.Fatalf("Failed to create records table: %v", err)
	}

	return db
}

func TestCreateTable(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	data.CreateTable()

	_, err := db.Query("SELECT * FROM activities")
	if err != nil {
		t.Fatalf("Activities table does not exist: %v", err)
	}

	_, err = db.Query("SELECT * FROM records")
	if err != nil {
		t.Fatalf("Records table does not exist: %v", err)
	}
}

func TestAddActivity(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	data.AddActivity("Running")

	// Query the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM activities WHERE name = ?", "Running").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query activities: %v", err)
	}

	if count != 1 {
		t.Fatalf("Expected 1 activity, got %d", count)
	}
}

func TestGetActivitiesString(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO activities (name) VALUES ('Swimming'), ('Cycling')")
	if err != nil {
		t.Fatalf("Failed to insert test activities: %v", err)
	}

	activities := data.GetActivitiesString()
	if len(activities) != 2 {
		t.Fatalf("Expected 2 activities, got %d", len(activities))
	}
}

func TestAddRecord(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	data.AddRecord("Running", "2024-03-20", 60)

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM records WHERE activity = ?", "Running").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query records: %v", err)
	}

	if count != 1 {
		t.Fatalf("Expected 1 record, got %d", count)
	}
}

func TestGetRecordsString(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO records (activity, date, timepassed) VALUES ('Running', '2024-03-20', 60)")
	if err != nil {
		t.Fatalf("Failed to insert test records: %v", err)
	}

	records := data.GetRecordsString()
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}
}

func TestDeleteActivity(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO activities (name) VALUES ('Yoga')")
	if err != nil {
		t.Fatalf("Failed to insert activity: %v", err)
	}

	data.DeleteActivity("Yoga")

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM activities WHERE name = ?", "Yoga").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query activities: %v", err)
	}

	if count != 0 {
		t.Fatalf("Expected 0 activities, got %d", count)
	}
}

func TestDeleteRecord(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()

	_, err := db.Exec("INSERT INTO records (activity, date, timepassed) VALUES ('Yoga', '2024-03-20', 45)")
	if err != nil {
		t.Fatalf("Failed to insert test record: %v", err)
	}

	data.DeleteRecord()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM records").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query records: %v", err)
	}

	if count != 0 {
		t.Fatalf("Expected 0 records, got %d", count)
	}
}
