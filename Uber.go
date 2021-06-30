package main

type (
	uberMoveType  func(uber *Uber, deltax, deltay int)
	ubergetCoord  func(uber *Uber) (int, int)
	uberDirection func(uber *Uber, deltax, deltay int)
	uberCheckMove func(uber *Uber) bool
	uberSetClient func(uber *Uber, client2 *passenger) bool
	Uber          struct {
		id, x, y  int
		avalaible bool
		world     *world // Just a pretty language ♥
		move      uberMoveType
		getCoord  ubergetCoord
		direction uberDirection
		client    *passenger
		checkMove uberCheckMove
		setClient uberSetClient
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
		id:        id,
		avalaible: true,
		x:         coordX,
		y:         coordY,
		world:     world,
		client:    nil,
		move: func(uber *Uber, deltax int, deltay int) {
			uber.x += deltax
			uber.y += deltay
			uber.world.Ubertraveled += 1
		},
		getCoord: func(uber *Uber) (int, int) {
			return uber.x, uber.y
		},
		direction: func(uber *Uber, deltax, deltay int) {
			lx := deltax
			ly := deltay
			switch {
			case uber.y > ly:
				uber.move(uber, 0, -1)
			case uber.y < ly:
				uber.move(uber, 0, 1)
			case uber.x > lx:
				uber.move(uber, -1, 0)
			case uber.x < lx:
				uber.move(uber, 1, 0)
			}
		},
		checkMove: func(uber *Uber) bool {
			if uber.client == nil || uber.avalaible {
				return false
			}
			if uber.client.x == uber.x && uber.client.y == uber.y && !uber.client.picked {
				uber.client.picked = true
			} else if uber.x == uber.client.objX && uber.y == uber.client.objY && !uber.client.done && uber.client.picked { // Done
				uber.client.done = true
				uber.client = nil
				uber.avalaible = true
				return false
			}
			if !uber.client.picked {
				uber.direction(uber, uber.client.x, uber.client.y)
			} else {
				uber.direction(uber, uber.client.objX, uber.client.objY)
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
				return true
			}
			return false
		},
	}
}
