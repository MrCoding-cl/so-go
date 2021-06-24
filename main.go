package main

import (
	"log"
	"time"
)

func main() {
	world := createWorld()
	addClientsToWorld(&world, "data\\requests_1.dat")
	addUbersToWorld(&world, "data\\manana.dat")
	start := time.Now()
	world.runwWithoutPram(&world)
	//world.runWithPram(&world)
	//world.runWithPram2Process(&world)
	end := time.Now()
	log.Println("Started Time: " + start.String())
	log.Println("End Time: " + end.String())
	log.Println("Total Time: ", end.Sub(start))
}
