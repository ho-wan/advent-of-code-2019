package main

import (
	"fmt"
	"math"
)

func main() {
	mass := []int{14}
	f := calcFuel(mass)
	fmt.Printf("Fuel: %d", f)
}

func calcFuel(mass []int) int {
	var f int
	for _, m := range mass {
		f += int(math.Max(math.Floor(float64(m)/3.0)-2.0, 0.0))
	}
	return f
}
