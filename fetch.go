package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import the driver
)

// return downloadLink, binPath, symlink,autoDownload
func Fetch(app string) (string, string, string, bool) {
	db, err := sql.Open("sqlite", "repos.db")
	if err != nil {
		log.Fatalf("Unable to open repo db %v\n", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT download_link, bin_path, symlink, auto_download FROM packages WHERE name = ?", app)
	if err != nil {
		log.Fatalf("Unable to query %v\n", err)
	}
	var download_link, bin_path, symlink string
	var auto_download bool

	for rows.Next() {

		err = rows.Scan(&download_link, &bin_path, &symlink, &auto_download)
		if err != nil {
			log.Fatalf("Unable to fetch some app info %v\n", err)
		}
	}
	return download_link, bin_path, symlink, auto_download

}
