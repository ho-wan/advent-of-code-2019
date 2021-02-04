package day07

import (
	"fmt"
	"strings"

	"github.com/ho-wan/advent-of-code-2019/2020/helpers"
)

const shinyGold = Bag("shiny gold")

// Bag describes type of bag
type Bag string

// Finder ...
type Finder struct {
	FoundBags map[Bag]bool
}

// Day07 ...
func Day07(filename string) int {
	input, err := helpers.ReadLines(filename)
	if err != nil {
		fmt.Println("faled to read input")
	}

	bags, err := inputToBags(input)
	if err != nil {
		fmt.Println("failed to convert input to bags")
	}

	f := &Finder{FoundBags: make(map[Bag]bool)}

	wantBags := []Bag{shinyGold}
	for len(wantBags) > 0 {
		wantBags = f.getBags(bags, wantBags)
	}

	return len(f.FoundBags)
}

func inputToBags(input []string) (map[Bag][]Bag, error) {
	bags := map[Bag][]Bag{}
	for _, line := range input {
		ss := strings.Split(line, "bags contain")
		bagKey := Bag(strings.Trim(ss[0], " "))

		contained := strings.Split(ss[1], ",")
		for _, bag := range contained {
			bagTrimmed := strings.Trim(bag, " ")
			if bagTrimmed == "no other bags." {
				bags[bagKey] = []Bag{}
				continue
			}

			split := strings.Split(bagTrimmed, " ")
			color := strings.Join(split[1:3], " ")

			bags[bagKey] = append(bags[bagKey], Bag(color))
		}
	}
	return bags, nil
}

func (f *Finder) getBags(bags map[Bag][]Bag, wantBags []Bag) []Bag {
	found := []Bag{}
	for k, subBags := range bags {
		for _, subBag := range subBags {
			for _, wb := range wantBags {
				if subBag == wb {
					found = append(found, k)
					f.FoundBags[k] = true
				}
			}
		}
	}
	return found
}
