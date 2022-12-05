package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	weights, err := parseInputAtoi("./input")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	idx := len(weights)
	fmt.Println(weights[idx-1])
	topThree := 0
	for i := 1; i < 4; i++ {
		topThree += weights[idx-i]
	}
	fmt.Println(topThree)
}

func parseInputAtoi(file string) ([]int, error) {
	out := make([]int, 0)
	sum := 0
	fileHandle, err := os.Open(file)
	if err != nil {
		return []int{}, err
	}

	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			out = append(out, sum)
			sum = 0
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			sum += value
		}
	}
	out = append(out, sum)
	sort.Ints(out)
	return out, nil
}

func parseInputScan(file string) ([]int, error) {
	out := make([]int, 0)
	sum := 0
	fileHandle, err := os.Open(file)
	if err != nil {
		return []int{}, err
	}

	s := bufio.NewScanner(fileHandle)
	var cur int
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			out = append(out, sum)
			sum = 0
		} else {
			_, err := fmt.Sscanf(line, "%d", &cur)
			if err != nil {
				log.Fatal(err)
			}
			sum += cur
		}
	}
	out = append(out, sum)
	sort.Ints(out)
	return out, nil
}
