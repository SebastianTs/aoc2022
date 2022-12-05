package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type rucksack struct {
	a map[rune]bool //compartment a
	b map[rune]bool //compartment b
}

const groupSize = 3

func main() {
	rs, err := parseInput("./input")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	fmt.Println(sumPriorities(rs))
	fmt.Println(sumPrioritiesBadges(rs))
}

func parseInput(file string) ([]rucksack, error) {
	out := make([]rucksack, 0)
	fileHandle, err := os.Open(file)
	if err != nil {
		return []rucksack{}, err
	}

	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		line := s.Text()
		l := len(line)
		if l%2 != 0 {
			return nil, errors.New("length of input string is odd, but must be even")
		} else {
			cur := rucksack{make(map[rune]bool), make(map[rune]bool)}
			for i, r := range line {
				if i < l/2 {
					cur.a[r] = true
				} else {
					cur.b[r] = true
				}
			}
			out = append(out, cur)
		}
	}
	if len(out)%groupSize != 0 {
		return nil, errors.New("length of input file (=#rucksack) does not match desired groupsize")
	}
	return out, nil
}

func findCommonItemInCompartment(r rucksack) (rune, error) {
	for c := range r.a {
		if _, exists := r.b[c]; exists {
			return c, nil
		}
	}
	return 0, errors.New("no common Item")
}

func calculatePriority(c rune) int {
	//Lowercase item types a through z have priorities 1 through 26.
	//Uppercase item types A through Z have priorities 27 through 52.
	v := int(c)
	if v > 96 {
		return v - 96
	}
	return v - 38
}

func sumPriorities(rs []rucksack) int {
	sum := 0
	for _, r := range rs {
		c, err := findCommonItemInCompartment(r)
		if err != nil {
			log.Panic(err)
		}
		sum += calculatePriority(c)
	}
	return sum
}

func findCommonItemInGroup(rs []rucksack) rune {
	cur := make(map[rune]int)
	sum := 0
	for _, r := range rs {
		for c := range r.a {
			cur[c]++
		}
		for c := range r.b {
			cur[c]++
		}
		sum++
	}

	for k, v := range cur {
		if v == sum {
			return k
		}
	}
	//Should never be reached
	return '🤕'
}

func sumPrioritiesBadges(rs []rucksack) int {
	sum := 0
	for i := 0; i < len(rs); i += groupSize {
		badge := findCommonItemInGroup(rs[i : i+groupSize])
		sum += calculatePriority(badge)
	}
	return sum
}
