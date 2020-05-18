package main

import "testing"

func TestMain(t *testing.T) {
	tt := []struct {
		name string
		w1   string
		w2   string
		expD int
	}{
		{"case1", "R8,U5,L5,D3", "U7,R6,D4,L4", 6},
		{"case2", "R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"case2", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := getDist(tc.w1, tc.w2)
			if d != tc.expD {
				t.Fatalf("got %v, want %v", d, tc.expD)
			}
		})
	}

}
