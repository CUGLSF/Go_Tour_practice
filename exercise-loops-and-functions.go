// exercise-loops-and-functions.go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		if math.Abs(z*z-x) > 0.01 {
			z = z - (z*z-x)/2*z
			fmt.Println(z)
		} else {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
