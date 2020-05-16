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
	errConv = errors.New("error converting string to int")
)

func readLines(path string) ([]string, error) {
	file, err := os.Open("day1-input.txt")
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
			return nil, errConv
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func calcFuel(mass []int) int {
	var f int
	for _, m := range mass {
		// divide mass by 3, minus 2 to get fuel - should not be less than zero
		ff := math.Floor(float64(m)/3.0) - 2.0
		f += int(math.Max(ff, 0.0))
	}
	return f
}

func main() {
	lines, err := readLines("day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mass, err := convertStringToInt(lines)
	if err != nil {
		log.Fatal(err)
	}

	f := calcFuel(mass)
	fmt.Printf("Fuel: %d", f)
}
