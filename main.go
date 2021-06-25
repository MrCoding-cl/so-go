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
	log.Println("MonoThread process")
	log.Println("Started Time:", start)
	log.Println("End Time:", end)
	log.Println("Total Time:", end.Sub(start))
	time.Sleep(5 * time.Second)
	start = time.Now()
	world = createWorld()
	addClientsToWorld(&world, "data\\requests_1.dat")
	addUbersToWorld(&world, "data\\manana.dat")
	//world.runwWithoutPram(&world)
	world.runWithPram(&world)
	//world.runWithPram2Process(&world)
	end = time.Now()
	log.Println("MultiThread process")
	log.Println("Started Time:", start)
	log.Println("End Time:", end)
	log.Println("Total Time:", end.Sub(start))
}
