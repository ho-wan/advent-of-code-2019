package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	errConv = errors.New("cannot convert string to int")
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func convertStringToInt(strings []string) ([]int, error) {
	nums := make([]int, 0, len(strings))
	for _, s := range strings {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("convertStringToInt: error: %w", errConv)
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func fuel(f float64) float64 {
	return math.Max(math.Floor(float64(f)/3.0)-2.0, 0.0)
}

func calcFuel(mass []int) int {
	var f int
	for _, m := range mass {
		// divide mass by 3, minus 2 to get fuel - should not be less than zero
		addFuel := fuel(float64(m))
		f += int(addFuel)
	}
	return f
}

func calcCumFuel(mass []int) int {
	var f int
	for _, m := range mass {
		var cf float64
		addFuel := float64(m)
		for {
			// divide mass by 3, minus 2 to get fuel, keep summing until addn mass is zero
			addFuel = fuel(float64(addFuel))
			if addFuel == 0 {
				break
			}
			cf += addFuel
		}
		f += int(cf)
	}
	return f
}

func main() {
	lines, err := readLines("day01-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mass, err := convertStringToInt(lines)
	if err != nil {
		log.Fatal(err)
	}

	f := calcFuel(mass)
	fmt.Printf("Fuel: %d\n", f)

	cf := calcCumFuel(mass)
	fmt.Printf("Cumulative Fuel: %d\n", cf)
}
