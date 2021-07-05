package main

import (
	"errors"
	"math"
	"math/rand"
	"path/filepath"
	"time"
)

func selectConditionedUber(ubers map[*Uber]float64) *Uber {
	rand.Seed(time.Now().Unix())
	acum := 0.0
	var uberWon *Uber
	for uberWon == nil {
		winnerprob := rand.Float64()
		for key, prob := range ubers {
			acum += prob
			if winnerprob <= acum {
				uberWon = key
				break
			}
		}
	}
	return uberWon
}

func DistanceBetween(client2 *passenger, uber *Uber) float64 {
	var deltaX, deltaY float64
	deltaX = math.Abs(float64(client2.x - uber.x))
	deltaY = math.Abs(float64(client2.y - uber.y))
	return deltaY + deltaX
}

func fileRoutine(world *world, requestPath string, timePath string) error {
	addClientsToWorld(world, requestPath)
	addUbersToWorld(world, timePath)
	var err error = nil
	if len(world.clients) == 0 || len(world.ubers) == 0 {
		err = errors.New("Clients or ubers doesn't exist")
	}
	return err
}

func morningRoutine(world *world) error {
	return fileRoutine(world, filepath.Join("./", "data", "requests_1.dat"), filepath.Join("./", "data", "manana.dat"))
}
func afternoonRoutine(world *world) error {
	return fileRoutine(world, filepath.Join("./", "data", "requests_2.dat"), filepath.Join("./", "data", "tarde.dat"))
}
func nightRoutine(world *world) error {
	return fileRoutine(world, filepath.Join("./", "data", "requests_3.dat"), filepath.Join("./", "data", "noche.dat"))
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

func getRoutine(client *Client) error {
	client.World = createWorld(12000)
	switch client.Config.RunType {
	case 0:
		err := morningRoutine(client.World)
		if err != nil {
			return err
		}
	case 1:
		err := afternoonRoutine(client.World)
		if err != nil {
			return err
		}
	case 2:
		err := nightRoutine(client.World)
		if err != nil {
			return err
		}
	case 3:
		randomRoutine(client.World)
	case 4:
		err := CustomRoutine(client.World, client)
		if err != nil {
			return err
		}
	default:
		return errors.New("RunType Not Found")
	}
	if client.Config.Pram {
		client.World.runWithPram(client.World)
	} else {
		client.World.runwWithoutPram(client.World)
	}
	return nil
}
