package main

type passenger struct {
	id, x, y, objX, objY, time int
	waiting, picked, done      bool
}

func createUberPassenger(id, x, y, objx, objy, time int, world2 *world) passenger {
	coordX := x
	coordY := y
	coordObjX := objx
	coordObjY := objy
	switch {
	case coordX > world2.maxX:
		coordX = world2.maxX
	case coordX < 0:
		coordX = 0
	}
	switch {
	case coordObjX > world2.maxX:
		coordObjX = world2.maxX
	case coordObjX < 0:
		coordObjX = 0
	}
	switch {
	case coordY > world2.maxY:
		coordY = world2.maxY
	case coordY < 0:
		coordY = 0
	}
	switch {
	case coordObjY > world2.maxY:
		coordObjY = world2.maxY
	case coordObjY < 0:
		coordObjY = 0
	}
	return passenger{
		id:      id,
		x:       coordX,
		y:       coordY,
		objX:    coordObjX,
		objY:    coordObjY,
		time:    time,
		waiting: true,
		picked:  false,
		done:    false,
	}
}
