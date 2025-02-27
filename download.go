package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Download(downloadLink string) {
	req, err := http.NewRequest("GET", downloadLink , nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer resp.Body.Close()

	// Check if ContentLength is available
	if resp.ContentLength == -1 {
		fmt.Println("Warning: ContentLength is not available, progress bar may not show.")
	}

	lastSlashIndex := strings.LastIndex(downloadLink, "/")
	fileName := downloadLink[lastSlashIndex+1:]

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)
	if err != nil {
		fmt.Println("Error copying data:", err)
	}
}
