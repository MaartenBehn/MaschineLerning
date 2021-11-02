package main

import "math"

func sigmoidFunc(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

func sigmoidDerivationFunc(x float64) float64 {
	return sigmoidFunc(x) * (1 - sigmoidFunc(x))
}
