package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntcodeLine(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		expRes []int
	}{
		{"case1", []int{1, 0, 0, 0, 99}, []int{4, 0, 0, 0, 99}},
		{"case2", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"case3", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"case4", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := runProgram(tc.input)
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !cmp.Equal(res, tc.expRes) {
				t.Fatalf("got %v, want %v", res, tc.expRes)
			}
		})
	}
}
