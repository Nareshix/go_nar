package main

import (
	"fmt"
	"os"
	"os/exec"
)


func DeleteApp(deleteLink string){
	cmd := exec.Command(os.Getenv("SHELL"), "-c", deleteLink)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Required if `wget` asks for input

	if err := cmd.Run(); err != nil {
		fmt.Println("Unable to delete app")
		return
	}
}
