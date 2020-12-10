package main

import (
	"bufio"
	"fmt"
	"os"
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

func day4() (int, error) {
	lines, err := readLines("input.txt")
	if err != nil {
		return 0, err
	}

	fmt.Print(lines[0])
	return -1, nil
}

func main() {
	_, err := day4()
	if err != nil {
		panic(err)
	}
}
