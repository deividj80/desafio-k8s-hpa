package main

import (
	"math"
)

func Calculation(x float64) float64 {
		x += math.Sqrt(x)
		return x
}