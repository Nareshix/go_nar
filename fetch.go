package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import the driver
)

// return downloadLink, binPath, symlink,deleteCmd,autoDownload
func Fetch(app string) (string, string, string, string, bool) {
	db, err := sql.Open("sqlite", "repos.db")
	if err != nil {
		log.Fatalf("Unable to open repo db %v\n", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT download_link, bin_path, symlink, delete_link, auto_download FROM packages WHERE name = ?", app)
	if err != nil {
		log.Fatalf("Unable to query %v\n", err)
	}
	var download_link, bin_path, symlink, delete_link string
	var auto_download bool

	for rows.Next() {

		err = rows.Scan(&download_link, &bin_path, &symlink, &delete_link, &auto_download)
		if err != nil {
			log.Fatalf("Unable to fetch some app info %v\n", err)
		}
	}
	return download_link, bin_path, symlink, delete_link, auto_download

}

// func Fetch(app string) (string,string,string,bool,string) {
// 	url := "https://raw.githubusercontent.com/Nareshix/repos/refs/heads/main/apps.json"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	downloadKey := app + ".download"
// 	binKey := app + ".bin_path"
// 	symlinkKey := app + ".symlink"
// 	autoDownloadKey := app + ".auto_download"
// 	deleteKey := app + ".delete"

// 	downloadLink := gjson.GetBytes(body, downloadKey)
// 	binPath := gjson.GetBytes(body, binKey)
// 	symlink := gjson.GetBytes(body, symlinkKey)
// 	autoDownload := gjson.GetBytes(body, autoDownloadKey)
// 	deleteCmd := gjson.GetBytes(body, deleteKey)

// 	if downloadLink.Exists(){
// 		return downloadLink.Str , binPath.Str, symlink.Str, autoDownload.Bool(),deleteCmd.Str
// 	}else{
// 		return "DownloadLink Missing", "BinPath missing", "Symlink missing", false,"No delete"
// 	}
// }
