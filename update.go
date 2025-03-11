package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import the driver
)

func UpdateDB(appName string, appVersion string) {

	url := "https://raw.githubusercontent.com/Nareshix/repo_changes/refs/heads/main/changes.txt"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	data := strings.Split(string(body), "\n")
	
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatalf("Unable to open data db %v\n", err)
	}
	defer db.Close()


	for _, cmd := range data{
		if cmd != ""{
			_, err := db.Exec(cmd)
			if err != nil{
				log.Fatal("unable to update db to latest")
			}
		}
	}
	
	if appName == ""{
		//TODO update all apps downloaded from a file after u update db
		fmt.Println("Successfully updated all")
		return
	}else if !upToDate(appName){
			deleteLink :=  FetchDel(appName)
			DeleteApp(deleteLink, appName)
	
			downloadLink, binPath, symlink,appVersion, autoDownload :=  Fetch(appName)
			Download(downloadLink, binPath, symlink,autoDownload, appName, appVersion)	
			fmt.Printf("Successfully updated %v\n", appName)
			return
		}
		fmt.Printf("%v is alr up-to-date\n", appName)
	}
	

	


func upToDate(appName string) bool{
	db, err := sql.Open("sqlite", "data.db")
    if err != nil {
        log.Fatal("Unable to open data db")
    }
    defer db.Close()


    
    row := db.QueryRow("SELECT version FROM downloaded WHERE name = ?", appName)

	var currentDownloadedVersion string
	err = row.Scan(&currentDownloadedVersion)
    if err != nil {
        log.Fatal("Unable to get user version for comparison")
    }
	
	if FetchLatestVersionNo(appName) == currentDownloadedVersion{
		return true
	}
	return false

}