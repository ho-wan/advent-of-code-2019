package main

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadlines(t *testing.T) {
	t.Run("hasFile", func(t *testing.T) {
		lines, err := readLines("day1-input.txt")
		if err != nil {
			t.Fatalf("error reading file")
		}
		if len(lines) == 0 {
			t.Fatalf("file should not be empty")
		}
	})

	t.Run("noFile", func(t *testing.T) {
		_, err := readLines("no-such-file.txt")
		if err == nil {
			t.Fatalf("did not get expected error")
		}
	})
}

func TestConvert(t *testing.T) {
	tt := []struct {
		name    string
		strings []string
		expInts []int
		expErr  error
	}{
		{"nil", nil, []int{}, nil},
		{"empty", []string{}, []int{}, nil},
		{"success", []string{"12", "23"}, []int{12, 23}, nil},
		{"notNumber", []string{"12", "ab"}, nil, errConv},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ints, err := convertStringToInt(tc.strings)
			if !errors.Is(err, tc.expErr) {
				t.Fatalf("error: got '%v', want '%v'", err, tc.expErr)
			}

			if !cmp.Equal(ints, tc.expInts) {
				t.Fatalf("got %v, want %v", ints, tc.expInts)
			}
		})
	}
}

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
