package main

import (
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

func DistanceBetween(client2 *client, uber *Uber) float64 {
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
