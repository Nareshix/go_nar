package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // Import the driver
)

func List(){
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatalf("Unable to open data db %v\n", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name FROM downloaded")
	if err != nil{
		log.Fatal("unable to read all downloaded programs ")
	}

	for rows.Next(){
		var name string
		err = rows.Scan(&name)
		if err != nil{
			log.Fatal("Unable to list all downloaded programs ")
		}
		fmt.Println(name)
	}
}
