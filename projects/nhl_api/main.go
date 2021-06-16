package main

import (
	"io"
	"log"
	"nhl_api/nhlApi"
	"os"
	"time"
)

func main() {
	// benchmarking the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {

		log.Fatalf("Error opening the file rosters.txt : %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()

	if err != nil {

		log.Fatalf("Error while getting all teams: %v", err)
	}

	for _, team := range teams {
		log.Println("----------------------------")
		log.Printf("Name: %v,\t City: %v", team.Name, team.Venue.City)
		log.Println("----------------------------")
	}

	log.Printf("took %v", time.Now().Sub(now).String())

}
