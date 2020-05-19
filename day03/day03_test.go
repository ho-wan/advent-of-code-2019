package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStrToVec(t *testing.T) {
	tt := []struct {
		name string
		in   string
		wOut [][2]int
	}{
		{"c1", "R8,U5,L5,D3", [][2]int{{8, 0}, {0, 5}, {-5, 0}, {0, -3}}},
		{"c2", "U7,R6,D4,L4", [][2]int{{0, 7}, {6, 0}, {0, -4}, {-4, 0}}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d, err := strToVec(tc.in)
			if err != nil {
				t.Fatalf("go unexpected error: %v", err)
			}
			if !cmp.Equal(d, tc.wOut) {
				t.Fatalf("got %v, want %v", d, tc.wOut)
			}
		})
	}

}

func TestGetCoords(t *testing.T) {
	tt := []struct {
		name string
		in   [][2]int
		wOut [][2]int
	}{
		{"c1", [][2]int{{2, 0}, {0, -1}}, [][2]int{{1, 0}, {2, 0}, {2, -1}}},
		{"c1", [][2]int{{1, 0}, {0, 1}, {-2, 0}}, [][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := getCoords(tc.in)
			if !cmp.Equal(out, tc.wOut) {
				t.Fatalf("got %v, want %v", out, tc.wOut)
			}
		})
	}
}

func TestGetDist(t *testing.T) {
	tt := []struct {
		name string
		w1   string
		w2   string
		wOut int
	}{
		{"c1", "R8,U5,L5,D3", "U7,R6,D4,L4", 6},
		{"c2", "R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"c3", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d, err := getDist(tc.w1, tc.w2)
			if err != nil {
				t.Fatalf("go unexpected error: %v", err)
			}
			if d != tc.wOut {
				t.Fatalf("got %v, want %v", d, tc.wOut)
			}
		})
	}

}
