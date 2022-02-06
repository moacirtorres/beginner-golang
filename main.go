package main

import (
	"fmt"
	"io"
	"log"
	"nhlApiDemo/nhlApi"
	"os"
	"time"
)

func main() {
	// help benchmarking the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("roster.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening the file roster.txt: %v", err)
	}

	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %v", err)
	}

	for _, team := range teams {
		fmt.Println("Name %s", team.Name)
	}

	log.Printf("took %v", time.Now().Sub(now).String())
}
