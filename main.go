package main

import (
	"github.com/Shea11012/go_utils/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}