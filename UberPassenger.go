package main

type passenger struct {
	id, x, y, objX, objY, time int
	waiting, picked, done      bool
	world                      *world
}

func createUberPassenger(id, x, y, objx, objy, time int, world2 *world) passenger {
	return passenger{
		id:      id,
		x:       x,
		y:       y,
		objX:    objx,
		objY:    objy,
		time:    time,
		waiting: true,
		picked:  false,
		done:    false,
		world:   world2,
	}
}
