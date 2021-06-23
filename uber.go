package main

import "log"

type uberMoveType func(uber *Uber, deltax, deltay int)
type ubergetCoord func(uber *Uber) (int, int)
type ubersetPathtoObjetive func(uber *Uber, deltax, deltay int)
type ubermakeMove func(uber *Uber) bool
type ubersetClient func(uber *Uber, client2 *client) bool

type Uber struct {
	id, x, y, movements, currentIndex int
	avalaible                         bool
	pathtoObj                         []rune
	world                             *world // Just a pretty language ♥
	possibles                         map[string]struct{ x, y int }
	move                              uberMoveType
	getCoord                          ubergetCoord
	setPathtoObjetive                 ubersetPathtoObjetive
	client                            *client
	makeMove                          ubermakeMove
	setClient                         ubersetClient
}

func createUber(id, x, y int, world *world) Uber {
	possibles := make(map[string]struct{ x, y int })
	possibles["S"] = struct{ x, y int }{x: 0, y: -1}
	possibles["N"] = struct{ x, y int }{x: 0, y: 1}
	possibles["E"] = struct{ x, y int }{x: 1, y: 0}
	possibles["W"] = struct{ x, y int }{x: -1, y: 0}
	return Uber{
		id:           id,
		avalaible:    true,
		x:            x,
		y:            y,
		movements:    0,
		currentIndex: 0,
		pathtoObj:    []rune(""), // Yup here is the STRING iterable of go c: ...
		world:        world,
		possibles:    possibles,
		client:       nil,
		move: func(uber *Uber, deltax int, deltay int) {
			uber.x += deltax
			uber.y += deltay
		},
		getCoord: func(uber *Uber) (int, int) {
			return uber.x, uber.y
		},
		setPathtoObjetive: func(uber *Uber, deltax, deltay int) {
			x := uber.x
			y := uber.y
			lx := deltax
			ly := deltay
			path := ""
			for !(x == lx && y == ly) {
				if y > ly {
					y -= 1
					path += "S"
				} else if y < ly {
					y += 1
					path += "N"
				}
				if x > lx {
					x -= 1
					path += "W"
				} else if x < lx {
					x += 1
					path += "E"
				}
			}
			uber.pathtoObj = []rune(path)
			uber.currentIndex = 0
		},
		makeMove: func(uber *Uber) bool {
			if len(uber.pathtoObj) == 0 || uber.client == nil || uber.avalaible {
				return false
			}
			if uber.client.x == uber.x && uber.client.y == uber.y && !uber.client.picked {
				uber.client.picked = true
				uber.setPathtoObjetive(uber, uber.client.objX, uber.client.objY)
				log.Println("Client Picked Uber ID", uber.id)
			} else if uber.x == uber.client.objX && uber.y == uber.client.objY && !uber.client.done { // Done
				uber.client.done = true
				uber.client = nil
				uber.pathtoObj = []rune("") // empty
				uber.currentIndex = 0
				uber.avalaible = true
				return false
			}
			letter := uber.pathtoObj[uber.currentIndex]
			move := uber.possibles[string(letter)]
			x, y := move.x, move.y
			if x == y { // 0,0
				return false
			}
			uber.move(uber, x, y)
			uber.currentIndex += 1
			uber.movements += 1
			uber.world.ubertraveled += 1
			return true
		},
		setClient: func(uber *Uber, client2 *client) bool {
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
				uber.setPathtoObjetive(uber, client2.x, client2.y)
				log.Println("Client received Uber ID", uber.id, "Distance to the client: ", DistanceBetween(client2, uber))
				return true
			}
			return false
		},
	}
}
