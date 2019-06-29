package main

import (
	"math/rand"
	"time"
)

var seed bool = false

// Returns a random integer between min and max
func randomRange(min, max int) int {
	if seed == false {
		rand.Seed(time.Now().UTC().UnixNano())
		seed = true
	}
	return rand.Intn(max-min) + min
}
