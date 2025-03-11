package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	versionB, err := os.ReadFile("ver.txt")
	if err != nil {
		log.Fatal()
	}
	// os.Stdout.Write(versionB)

	versionNo, err := strconv.Atoi(string(versionB))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get("https://raw.githubusercontent.com/Nareshix/repo_changes/refs/heads/main/data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("HTTP status not OK")
	}

	var data map[string][]string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	changedToLatest := false
	var values []string
	for k, tmpValue := range data {
		key_on_cloud, err := strconv.Atoi(k)
		if err != nil {
			log.Fatal(err)
		}
		if key_on_cloud > versionNo {
			versionNo = key_on_cloud
			changedToLatest = true
			values = tmpValue
		}
	}
	fmt.Println(values) //REMOVE
	if changedToLatest {
		err = os.WriteFile("ver.txt", []byte(strconv.Itoa(versionNo)), 0644)
		if err != nil {
			log.Fatal(err)
		}
		//TODO updates all the necessary packages
	}
	// converts the int -> str -> bin
}
