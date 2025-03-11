package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	_ "modernc.org/sqlite" // Import the driver
)

// checks for download script --> deb --> tar.gz

func Download(downloadLink string, binPath string, symlink string, autoDownload bool, appName string, appVersion string) {
	// If autodownload script is avail use that only
	if autoDownload {
		cmd := exec.Command(os.Getenv("SHELL"), "-c", downloadLink)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin // Required if `wget` asks for input

		if err := cmd.Run(); err != nil {
			fmt.Println("Unable to run autodownload script")
			return
		}
		StoreValueInTrackingDB(appName, appVersion)
		return
	}

	// grabs the tar.gz or .deb file from github via http
	cmd := exec.Command("sudo", "wget", "-q", "--show-progress", downloadLink)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Required if `wget` asks for input

	if err := cmd.Run(); err != nil {
		fmt.Println("Unable to download http link")
		return
	}

	parsedDownloadLink, err := url.Parse(downloadLink)
	if err != nil {
		fmt.Println("Unable to parse download link")
		return
	}
	fileNameWithExtension := path.Base(parsedDownloadLink.Path)

	// Donwload DEB file
	if strings.HasSuffix(fileNameWithExtension, ".deb") {

		debFileName := "./" + fileNameWithExtension
		cmd := exec.Command("sudo", "apt", "install", debFileName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin // Required if `wget` asks for input

		if err := cmd.Run(); err != nil {
			fmt.Println("Unable to install DEB file")
			return
		}

		// deleted the .deb file (not needed)
		if err := os.Remove(fileNameWithExtension); err != nil {
			fmt.Println("Unable to remove DEB file")
			return
		}
		StoreValueInTrackingDB(appName, appVersion)
		return

		// download tar binaries in /opt/ and create a symlink to /usr/local/bin
	} else if strings.HasSuffix(fileNameWithExtension, ".tar.gz") {

		cmd1 := exec.Command("sudo", "tar", "-xvzf", fileNameWithExtension, "-C", "/opt")
		if err := cmd1.Run(); err != nil {
			fmt.Println("failed to untar file")
		}

		// remove the tar.gz file (not needed)
		if err := os.Remove(fileNameWithExtension); err != nil {
			fmt.Println("Unable to remove tar.gz. file")
			return
		}

		fullPathToBin := "/opt/" + binPath
		fullPathToSymlink := "/usr/local/bin/" + symlink
		cmd2 := exec.Command("sudo", "ln", "-s", fullPathToBin, fullPathToSymlink)
		if err := cmd2.Run(); err != nil {
			fmt.Println("failed to create symlink")
			return

		}
		StoreValueInTrackingDB(appName, appVersion)
		return
	}
}

//TODO handles case when alr downloaded
func StoreValueInTrackingDB(appName string, version string){
	db, err := sql.Open("sqlite", "data.db")
    if err != nil {
        log.Fatal("Unable to open data db")
    }
    defer db.Close()

    
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS downloaded (name TEXT PRIMARY KEY, version TEXT )")
    if err != nil {
        log.Fatal("Downloaded table does not exist and unable to create one")
    }

    _, err = db.Exec("INSERT INTO downloaded (name, version) VALUES (?, ?)", appName, version)
    if err != nil {
        log.Fatal("Unable to update downloaded tables in db")
    }
}