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

var errConv = errors.New("cannot convert string to int")

func readNums(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var nums []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		strNums := strings.Split(line, ",")
		for _, str := range strNums {
			si, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("convertStringToInt: error: %w", errConv)
			}
			nums = append(nums, si)
		}
	}
	return nums, nil
}

func runProgram(in []int) ([]int, error) {
	nums := in
	for i := 0; i < len(nums); i += 4 {
		switch nums[i] {
		case 99:
			break
		case 1:
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		case 2:
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
		}
	}
	return nums, nil
}

func main() {
	nums, err := readNums("day02-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res, err := runProgram(nums)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %v", res[0])
}
