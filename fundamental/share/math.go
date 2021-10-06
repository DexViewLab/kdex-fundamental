package share

import (
	"math/rand"
	"time"
)

var (
	// EPSILON ..
	EPSILON = 0.00000001
	// R ..
	R = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandFloat ..
func RandFloat(min, max float64) float64 {
	return min + R.Float64() * (max - min)
}

// Max ..
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min ..
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// FloatEquals ..
func FloatEquals(a, b float64) bool {
	if (a - b) < EPSILON && (b - a) < EPSILON {
		return true
	}
	return false
}