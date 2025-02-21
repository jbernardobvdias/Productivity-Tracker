package data

import (
	"database/sql"
	"fmt"
	"log"
)

func LoadTable() {
	db, err := sql.Open("postgress", "creds")

	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM")

	if err != nil {
		log.Fatal("There was a problem with the querry.")
	}

	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, name)
	}
}
