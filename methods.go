package main

import (
	"math"
	"math/rand"
	"time"
)

func selectConditionedUber(ubers map[*Uber]float64, c int) *Uber { // Conditioned random select...
	rand.Seed(time.Now().Unix())
	newmap := make(map[*Uber]float64)
	for len(newmap) == 0 {
		selector := rand.Float64()
		for uber, f := range ubers {
			if f > selector {
				newmap[uber] = f
			}
		}
	}
	if c == 10 { // Limit. because when it is reached all ubers haves similar probs
		winner := rand.Intn(len(newmap))
		c := 0
		for uber, _ := range newmap {
			if c == winner {
				return uber
			}
			c += 1
		}
	}
	if len(newmap) > 2 {
		total := 0.0
		for _, f := range newmap {
			total += f
		}
		for uber, f := range newmap {
			newmap[uber] = f / total // Makes a new prob with the winners
		}
		return selectConditionedUber(newmap, c+1)
	} else if len(newmap) == 2 {
		total := 0.0
		for _, f := range newmap {
			if total != 0 {
				if total == f { // the probability is small but not 0
					result := rand.Intn(2) // 50 to 50.
					c := 0
					for uber := range newmap {
						if c == result {
							return uber
						}
						c += 1
					}
				}
			}
			total += f
		}
		for uber, f := range newmap {
			newmap[uber] = f / total
		}
		return selectConditionedUber(newmap, c+1)
	}
	var result *Uber = nil
	for uber := range newmap {
		result = uber
	}
	return result
}

func DistanceBetween(client2 client, uber Uber) float64 {
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
