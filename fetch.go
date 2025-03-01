package main

import (
	// "fmt"
	// "fmt"
	// "fmt"
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func Fetch(app string) (string,string,string,bool) {
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

	downloadKey := app + ".download"
	binKey := app + ".bin_path"
	symlinkKey := app + ".symlink"
	autoDownloadKey := app + ".auto_download"


	downloadLink := gjson.GetBytes(body, downloadKey)
	binPath := gjson.GetBytes(body, binKey)
	symlink := gjson.GetBytes(body, symlinkKey)
	autoDownload := gjson.GetBytes(body, autoDownloadKey)

	if downloadLink.Exists(){
		return downloadLink.Str , binPath.Str, symlink.Str, autoDownload.Bool()
	}else{
		return "DownloadLink Missing", "BinPath missing", "Symlink missing", false
	}
}
