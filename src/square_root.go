package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	p := 0.01

	for z * z < x {
		z++
	}
	z--

	for z * z < x {
		z += p
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
