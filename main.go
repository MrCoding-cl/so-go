package main

import (
	"log"
	"math/rand"
	"time"
)

func simplefunc(s *string, ch string) {
	*s = ch
}

func main() {
	world := createWorld()
	start := time.Now()
	for i := 0; i < 200; i++ {
		uber := createUber(i, rand.Intn(1000), rand.Intn(1000), &world)
		world.ubers = append(world.ubers, &uber)
	}
	for i := 0; i < 1000; i++ {
		client := createClient(i, rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(1000), rand.Intn(12000), &world)
		world.addClient(&world, &client)
	}
	world.runwWithoutPram(&world)
	end := time.Now()
	log.Println("Started Time: " + start.String())
	log.Println("End Time: " + end.String())
}
