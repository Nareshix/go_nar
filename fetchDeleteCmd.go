package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import the driver
)

func FetchDel(app string) (string){
	db, err := sql.Open("sqlite", "repos.db")
	if err != nil {
		log.Fatalf("Unable to open repo db %v\n", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT delete_link FROM packages WHERE name = ?", app)
	if err != nil {
		log.Fatalf("Unable to query %v\n", err)
	}

	var delete_link string
	for rows.Next() {
		err = rows.Scan(&delete_link)
		if err != nil {
			log.Fatalf("Unable to fetch delete link info %v\n", err)
		}
	}
	return delete_link

}
