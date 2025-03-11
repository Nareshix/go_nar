package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"database/sql"

	_ "modernc.org/sqlite"
)


func DeleteApp(deleteLink string, appName string){
	cmd := exec.Command(os.Getenv("SHELL"), "-c", deleteLink)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Required if `wget` asks for input

	if err := cmd.Run(); err != nil {
		log.Fatal("Unable to delete app")
		return
	}
	fmt.Printf("Succesfully deleted %v\n", appName )
	removeTrackingFromDB(appName)
	
}

func removeTrackingFromDB(appName string){
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatalf("Unable to open data db %v\n", err)
	}
	defer db.Close()
	
	_, err = db.Exec("DELETE FROM downloaded WHERE name = ?", appName)
    if err != nil {
        log.Fatal("Downloaded table does not exist and unable to create one")
    }
}