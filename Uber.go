package main

import "log"

type (
	uberMoveType          func(uber *Uber, deltax, deltay int)
	ubergetCoord          func(uber *Uber) (int, int)
	ubersetPathtoObjetive func(uber *Uber, deltax, deltay int)
	ubermakeMove          func(uber *Uber) bool
	ubersetClient         func(uber *Uber, client2 *passenger) bool
	Uber                  struct {
		id, x, y, movements, MovementsToReach int
		avalaible                             bool
		world                                 *world // Just a pretty language ♥
		move                                  uberMoveType
		getCoord                              ubergetCoord
		walktoObjetive                        ubersetPathtoObjetive
		client                                *passenger
		makeMove                              ubermakeMove
		setClient                             ubersetClient
	}
)

func createUber(id, x, y int, world *world) Uber {
	coordX := x
	coordY := y
	switch {
	case coordX > world.maxX:
		coordX = world.maxX
	case coordX < 0:
		coordX = 0
	}
	switch {
	case coordY > world.maxY:
		coordY = world.maxY
	case coordY < 0:
		coordY = 0
	}
	return Uber{
		id:               id,
		avalaible:        true,
		x:                coordX,
		y:                coordY,
		movements:        0,
		MovementsToReach: 0,
		world:            world,
		client:           nil,
		move: func(uber *Uber, deltax int, deltay int) {
			uber.x += deltax
			uber.y += deltay
			uber.movements += 1
			uber.world.Ubertraveled += 1
			uber.MovementsToReach += 1
		},
		getCoord: func(uber *Uber) (int, int) {
			return uber.x, uber.y
		},
		walktoObjetive: func(uber *Uber, deltax, deltay int) {
			lx := deltax
			ly := deltay
			if uber.y > ly {
				uber.move(uber, 0, -1)
			} else if uber.y < ly {
				uber.move(uber, 0, 1)
			} else if uber.x > lx {
				uber.move(uber, -1, 0)
			} else if uber.x < lx {
				uber.move(uber, 1, 0)
			}
		},
		makeMove: func(uber *Uber) bool {
			if uber.client == nil || uber.avalaible {
				return false
			}
			if uber.client.x == uber.x && uber.client.y == uber.y && !uber.client.picked {
				uber.client.picked = true
				log.Println("Uber", uber.id, "Picked passenger in", uber.MovementsToReach, "Movements")
				uber.MovementsToReach = 0
			} else if uber.x == uber.client.objX && uber.y == uber.client.objY && !uber.client.done && uber.client.picked { // Done
				log.Println("Uber", uber.id, "DONE in", uber.MovementsToReach, "Its ok?")
				if DistanceBetweenPoints(uber.client.x, uber.client.y, uber.client.objX, uber.client.objY) == uber.MovementsToReach {
					log.Println("YUP")
				} else {
					realmovements := DistanceBetweenPoints(uber.client.x, uber.client.y, uber.client.objX, uber.client.objY)
					log.Println("NOPE MAX Movement are:", realmovements)
				}
				uber.client.done = true
				uber.client = nil
				uber.MovementsToReach = 0
				uber.avalaible = true
				return false
			}
			if !uber.client.picked {
				uber.walktoObjetive(uber, uber.client.x, uber.client.y)
			} else {
				uber.walktoObjetive(uber, uber.client.objX, uber.client.objY)
			}
			return true
		},
		setClient: func(uber *Uber, client2 *passenger) bool {
			if uber.client == nil {
				uber.client = client2
			} else {
				if uber.client.done {
					uber.client = client2
				}
			}
			if uber.client == client2 {
				uber.client.waiting = false
				uber.avalaible = false
				uber.walktoObjetive(uber, client2.x, client2.y)
				return true
			}
			return false
		},
	}
}
