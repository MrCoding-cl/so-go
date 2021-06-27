package main

import (
	"errors"
	"math"
	"math/rand"
	"sort"
	"time"
)

func selectConditionedUber(ubers map[*Uber]float64) *Uber {
	/*
		Conditioned random select
		receives a map of the ubers with their respective probabilities, sorts them from highest to lowest probability,
		then creates a summation of each element's probability with the previous one.
		The last element has a probability of 1. It draws a floating point number between 0 and 1 and starts checking
		from the first uber if it has a probability greater than that of the floating point and is selected,
		thus achieving uniformity and maintaining the rule of the probability difference between each of the elements.
	*/
	rand.Seed(time.Now().Unix())
	keys := make([]*Uber, 0)
	for key := range ubers {
		if ubers[key] > 0.0 {
			keys = append(keys, key)
		}
	}
	sort.Slice(keys, func(i, j int) bool { return ubers[keys[i]] > ubers[keys[j]] })
	acum := 0.0
	for _, key := range keys {
		ubers[key] += acum
		acum = ubers[key]
	}
	var uberWon *Uber = nil
	for uberWon == nil {
		winnerprob := rand.Float64()
		for _, key := range keys {
			uberprob := ubers[key]
			if winnerprob <= uberprob {
				uberWon = key
				break
			}
		}
	}
	return uberWon
}

func DistanceBetween(client2 *passenger, uber *Uber) float64 {
	var deltaX, deltaY float64
	if client2.waiting || !client2.picked {
		deltaX = math.Abs(float64(client2.x - uber.x))
		deltaY = math.Abs(float64(client2.y - uber.y))
	} else {
		deltaX = math.Abs(float64(client2.objX - uber.x))
		deltaY = math.Abs(float64(client2.objY - uber.y))
	}
	return deltaY + deltaX
}

func DistanceBetweenPoints(x1, y1, x2, y2 int) int {
	deltaX := math.Abs(float64(x1 - x2))
	deltaY := math.Abs(float64(y1 - y2))
	return int(deltaY + deltaX)
}

func fileRoutine(world *world, requestPath string, TimePath string) error {
	addClientsToWorld(world, requestPath)
	addUbersToWorld(world, TimePath)
	var err error = nil
	if len(world.clients) == 0 || len(world.ubers) == 0 {
		err = errors.New("Clients or ubers doesn't exist")
	}
	return err
}

func morningRoutine(world *world) error {
	return fileRoutine(world, "data\\requests_1.dat", "data\\manana.dat")
}
func afternoonRoutine(world *world) error {
	return fileRoutine(world, "data\\requests_2.dat", "data\\tarde.dat")
}
func nightRoutine(world *world) error {
	return fileRoutine(world, "data\\requests_3.dat", "data\\noche.dat")
}
func randomRoutine(world *world) {
	rand.Seed(time.Now().Unix())
	ubers := 300 - rand.Intn(200)     //  Min 100, Max 300
	clients := 3000 - rand.Intn(2000) // Min 1000, max 3000
	for i := 0; i < clients; i++ {
		client := createUberPassenger(i, rand.Intn(1000), rand.Intn(1000), rand.Intn(1000),
			rand.Intn(1000), rand.Intn(12000), world)
		world.addClient(world, &client)
	}
	for i := 0; i < ubers; i++ {
		uber := createUber(i, rand.Intn(1000), rand.Intn(1000), world)
		world.ubers = append(world.ubers, &uber)
	}
}
func CustomRoutine(world *world, client *Client) error {
	config := &client.Config
	for i, uber := range config.Ubers {
		ub := createUber(i, uber.X, uber.Y, world)
		world.ubers = append(world.ubers, &ub)
	}
	for i, passenger := range config.Request {
		pass := createUberPassenger(i, passenger.Xi, passenger.Yi, passenger.Xf, passenger.Yf, passenger.T, world)
		world.addClient(world, &pass)
	}
	var err error
	if len(world.clients) == 0 || len(world.ubers) == 0 {
		err = errors.New("Clients or ubers doesn't exist")
	}
	return err
}
