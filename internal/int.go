package internal

import (
	"math/rand"
)

func RandomInt(min, max int) int {
	return rand.Intn((max - min) + min)
}
