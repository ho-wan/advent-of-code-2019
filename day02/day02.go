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
	for i := 0; i <= len(nums); i += 4 {
		if nums[i] == 99 {
			break
		}
		if i > len(nums)-4 {
			return nums, errors.New("no solution, end of nums reached")
		}
		if nums[i] == 1 {
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		}
		if nums[i] == 2 {
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
	nums[1] = 12
	nums[2] = 2

	res, err := runProgram(nums)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %v\n", res[0])

	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			nums, err := readNums("day02-input.txt")
			if err != nil {
				log.Fatal(err)
			}

			nums[1] = i
			nums[2] = j
			res, err := runProgram(nums)
			if err != nil {
				fmt.Printf("error %v\n", err)
			}
			if res[0] == 19690720 {
				fmt.Printf("part2: %v\n", 100*i+j)
			}
		}
	}
}
