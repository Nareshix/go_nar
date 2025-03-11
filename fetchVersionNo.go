package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import the driver
)

func FetchLatestVersionNo(appName string) (string){
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatalf("Unable to open data db %v\n", err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT version FROM packages WHERE name = ?", appName)

	var version string
	err = row.Scan(&version)
	if err != nil{
		log.Fatal("Unable to fetch latest version no")
	}
	return version

}

func FetchCurrentVersionNo(appName string) (string){
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatalf("Unable to open data db %v\n", err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT version FROM downloaded WHERE name = ?", appName)

	var version string
	err = row.Scan(&version)
	if err != nil{
		log.Fatal("Unable to fetch latest version no")
	}
	return version

}