package main

import (
	"log"
	"sync"
)

type (
	world struct {
		maxX, maxY, time, ubertraveled int
		ubers                          []*Uber // ah... Pointers, Sweet XD
		clients                        map[int][]*client
		waitingclients                 []*client
		filterWaitingClients           worldFilterWaitingClients
		addClient                      worldAddClient
		clientsToWaitingList           worldclientstowaitinglist
		uberForClient                  worlduberforclient
		getAvalaibleUbers              worldGetAvalaibleUbers
		runwWithoutPram                worldRunWithoutPram
		runWithPram                    worldRunWithPram
		runWithPram2Process            worldRunWithPramTwoProcess
	}
	worldFilterWaitingClients  func(world2 *world)
	worldAddClient             func(world2 *world, client2 *client)
	worldclientstowaitinglist  func(world2 *world)
	worlduberforclient         func(world2 *world, client2 *client, ubers *[]*Uber) bool
	worldGetAvalaibleUbers     func(world2 *world) []*Uber
	worldRunWithoutPram        func(world2 *world)
	worldRunWithPram           func(world2 *world)
	worldRunWithPramTwoProcess func(world2 *world)
)

func createWorld() world {
	return world{
		maxX:           1000,
		maxY:           1000,
		time:           0,
		ubertraveled:   0,
		ubers:          make([]*Uber, 0),
		clients:        make(map[int][]*client),
		waitingclients: make([]*client, 0),
		filterWaitingClients: func(world2 *world) {
			oldlist := world2.waitingclients
			newlist := make([]*client, 0)
			for _, client := range oldlist {
				if client.waiting {
					newlist = append(newlist, client)
				}
			}
			world2.waitingclients = newlist
		},
		addClient: func(world2 *world, client2 *client) {
			if client2.time < world2.time {
				client2.time = world2.time
			}
			world2.clients[client2.time] = append(world2.clients[client2.time], client2)
		},
		clientsToWaitingList: func(world2 *world) {
			world2.waitingclients = append(world2.waitingclients, world2.clients[world2.time]...)
		},
		getAvalaibleUbers: func(world2 *world) []*Uber {
			oldList := world2.ubers
			newlist := make([]*Uber, 0)
			for _, uber := range oldList {
				if uber.avalaible {
					newlist = append(newlist, uber)
				}
			}
			return newlist
		},
		uberForClient: func(world2 *world, client2 *client, ubers *[]*Uber) bool {
			total := 0.0
			//ubers :=
			if len(*ubers) == 0 { // No ubers avalaible
				return false
			}
			for _, uber := range *ubers {
				total += DistanceBetween(client2, uber)
			}
			probs := make(map[*Uber]float64)
			for _, uber := range *ubers {
				probs[uber] = DistanceBetween(client2, uber) / total
			}
			uber := selectConditionedUber(probs)
			//if !uber.setClient(uber, client2) {
			//	world2.uberForClient(world2, client2)
			//}
			uber.setClient(uber, client2)
			return true
		},
		runwWithoutPram: func(world2 *world) {
			for world2.time < 12000 {
				world2.filterWaitingClients(world2)
				world2.clientsToWaitingList(world2)
				for _, client := range world2.waitingclients {
					ubers := world2.getAvalaibleUbers(world2)
					if len(ubers) == 0 {
						break
					}
					if !world2.uberForClient(world2, client, &ubers) {
						break // Because there's no more ubers avalaible
					}
				}
				for _, uber := range world2.ubers {
					uber.makeMove(uber)
				}
				world2.time += 1
				log.Println(world2.time, world2.ubertraveled)
			}
		},
		runWithPram: func(world2 *world) {
			for world2.time < 12000 {
				world2.filterWaitingClients(world2)
				world2.clientsToWaitingList(world2)
				var wg = new(sync.WaitGroup)
				for _, client := range world2.waitingclients {
					ubers := world2.getAvalaibleUbers(world2)
					if len(ubers) == 0 {
						break
					}
					wg.Add(1)
					client := client
					go func() {
						defer wg.Done()
						world2.uberForClient(world2, client, &ubers)
					}()
				}
				wg.Wait()
				for _, uber := range world2.ubers {
					if uber.client == nil {
						continue
					}
					wg.Add(1)
					go func(uber *Uber) {
						defer wg.Done()
						uber.makeMove(uber)
					}(uber)
					//uber.makeMove(uber)
				}
				wg.Wait()
				world2.time += 1
				log.Println(world2.time, world2.ubertraveled)
			}
		},
		runWithPram2Process: func(world2 *world) {
			for world2.time < 12000 {
				var wg = new(sync.WaitGroup)
				wg.Add(2)
				go func() {
					defer wg.Done()
					world2.filterWaitingClients(world2)
					world2.clientsToWaitingList(world2)
					for _, client := range world2.waitingclients {
						ubers := world2.getAvalaibleUbers(world2)
						if len(ubers) == 0 {
							break
						}
						if !world2.uberForClient(world2, client, &ubers) {
							break // Because there's no more ubers avalaible
						}
					}
				}()
				go func() {
					defer wg.Done()
					for _, uber := range world2.ubers {
						uber.makeMove(uber)
					}
				}()
				wg.Wait()
				world2.time += 1
				log.Println(world2.time, world2.ubertraveled)
			}
		},
	}
}
