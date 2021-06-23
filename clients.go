package main

type client struct {
	id, x, y, objX, objY, time int
	waiting, picked, done      bool
	world                      *world
}

func createClient(id, x, y, objx, objy, time int, world2 *world) client {
	return client{
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
