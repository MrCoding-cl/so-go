package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	world := createWorld()
	for i := 0; i < 200; i++ {
		uber := createUber(i, rand.Intn(1000), rand.Intn(1000), &world)
		world.ubers = append(world.ubers, &uber)
	}
	for i := 0; i < 500; i++ {
		client := createClient(i, rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000),
			rand.Intn(12000), &world)
		world.addClient(&world, &client)
	}
	start := time.Now()
	//world.runwWithoutPram(&world)
	world.runWithPram(&world)
	//world.runWithPram2Process(&world)
	end := time.Now()
	log.Println("Started Time: " + start.String())
	log.Println("End Time: " + end.String())
}
