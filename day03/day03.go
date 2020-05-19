package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Psuedocode
- convert (w1,w2) to slice of strings
- convert string to vector of dist, eg L2 = (-2, 0)
- for w1, start at (0,0), add each point to a slice of coords
- do the same for w2, if coord already in slice from w1, add to array "crossed"
- get manhattan dist from each cross to origin, and return smaller
*/

var errConv = errors.New("cannot convert string to int")

func strToVec(w1 string) ([][2]int, error) {
	sw1 := strings.Split(w1, ",")
	v := [][2]int{}
	for _, s := range sw1 {
		dir := s[:1]
		n := s[1:]
		dst, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("convertStringToInt: error: %w", errConv)
		}

		tv := [2]int{0, 0}
		switch dir {
		case "L":
			tv[0] -= dst
		case "R":
			tv[0] += dst
		case "D":
			tv[1] -= dst
		case "U":
			tv[1] += dst
		}
		v = append(v, tv)
	}
	return v, nil
}

func getCoords(co [][2]int) [][2]int {
	coords := [][2]int{}
	pos := [2]int{0, 0}
	for _, c := range co {
		if c[0] == 0 && c[1] == 0 {
			continue
		}
		for ; c[0] > 0; c[0]-- {
			pos[0]++
			coords = append(coords, pos)
		}
		for ; c[0] < 0; c[0]++ {
			pos[0]--
			coords = append(coords, pos)
		}
		for ; c[1] > 0; c[1]-- {
			pos[1]++
			coords = append(coords, pos)
		}
		for ; c[1] < 0; c[1]++ {
			pos[1]--
			coords = append(coords, pos)
		}
	}
	return coords
}

func getDist(w1, w2 string) (int, error) {
	vw1, err := strToVec(w1)
	if err != nil {
		return -1, fmt.Errorf("getDist: %w", err)
	}
	vw2, err := strToVec(w2)
	if err != nil {
		return -1, fmt.Errorf("getDist: %w", err)
	}

	cw1 := getCoords(vw1)
	cw2 := getCoords(vw2)

	overlap := [][2]int{}
	for i, c1 := range cw1 {
		for _, c2 := range cw2 {
			if c1[0] == c2[0] && c1[1] == c2[1] {
				overlap = append(overlap, c1)
				continue
			}
		}
		// for visual progress
		if i%10000 == 0 {
			fmt.Println(i)
		}
	}

	var minMd int
	for _, o := range overlap {
		md := absInt(o[0]) + absInt(o[1])
		if md < minMd || minMd == 0 {
			minMd = md
		}
	}
	return minMd, nil
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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

func main() {
	lines, err := readLines("day03-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got lines, length %d, %d", len(lines[0]), len(lines[1]))

	p1, err := getDist(lines[0], lines[1])
	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}
	fmt.Println("Part1: ", p1)
}
