package main

import (
	// "fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func Fetch(app string) string {
	url := "https://raw.githubusercontent.com/Nareshix/repos/refs/heads/main/apps.json"


	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// app.download
	var downloadKey strings.Builder 
	downloadKey.WriteString(app)
	downloadKey.WriteString(".download")

	downloadLink := gjson.GetBytes(body, downloadKey.String())

	if downloadLink.Exists(){
		return downloadLink.Str
	}else{
		return "no link"	
	}
}
