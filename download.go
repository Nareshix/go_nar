package main

import (
	// "encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	// "fmt"
)

func Download(downloadLink string, binPath string, symlink string, autoDownload bool) {
	
	// If autodownload script is avail use that only
	if autoDownload{
	fmt.Println(downloadLink)
	cmd := exec.Command(os.Getenv("SHELL"), "-c", downloadLink)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Required if `wget` asks for input
	if err := cmd.Run(); err != nil {
		fmt.Println("Unable to run autodownload script")
		return
	}
		return //exit program 
	}

	// gets the download link from http
	cmd := exec.Command("sudo", "wget","-q", "--show-progress", downloadLink)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Required if `wget` asks for input
	if err := cmd.Run(); err != nil {
		fmt.Println("Unable to download http link")
		return
	}

	parsedDownloadLink,err := url.Parse(downloadLink)
	if err != nil{
		fmt.Println("Unable to parse download link")
		return
	}
	fileNameWithExtension := path.Base(parsedDownloadLink.Path)

	// Donwload DEB file
	if strings.HasSuffix(fileNameWithExtension, ".deb"){

		debFileName := "./" + fileNameWithExtension
		cmd := exec.Command("sudo", "apt", "install", debFileName )
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin // Required if `wget` asks for input
		if err := cmd.Run(); err != nil{
			fmt.Println("Unable to install DEB file")
			return
		}

		if err := os.Remove(fileNameWithExtension); err != nil{
			fmt.Println("Unable to remove DEB file")
			return
		}
		
	// download tar binaries in /opt/ and create a symlink to /usr/local/bin
	} else if strings.HasSuffix(fileNameWithExtension, ".tar.gz"){
		
		cmd1 := exec.Command("sudo", "tar", "-xvzf", fileNameWithExtension,"-C", "/opt")
		if err := cmd1.Run(); err != nil{
			fmt.Println("failed to untar file")
		}
		
		if err := os.Remove(fileNameWithExtension); err != nil{
			fmt.Println("Unable to remove DEB file")
			return
		}
		
		fullPathToBin := "/opt/" + binPath
		cmd2 := exec.Command("sudo", "ln", "-s", fullPathToBin, symlink )
		if err := cmd2.Run(); err != nil{
			fmt.Println("failed to create symlink")
			return

		
	} 


	}
}