package main

import "testing"

func TestCalcFuel(t *testing.T) {
	tt := []struct {
		name string
		mass []int
		fuel int
	}{
		{"nil", nil, 0},
		{"zero", []int{0}, 0},
		{"case1", []int{14}, 2},
		{"case2", []int{1969}, 654},
		{"case3", []int{100756}, 33583},
		{"sum", []int{14, 1969, 100756}, 34239},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := calcFuel(tc.mass)
			if f != tc.fuel {
				t.Fatalf("expect fuel to be %d, got %d", tc.fuel, f)
			}
		})

	}
}
