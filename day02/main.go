package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type guide struct {
	l rune
	r rune
}

var score map[rune]int = map[rune]int{
	'A': 1, 'B': 2, 'C': 3,
	'X': 1, 'Y': 2, 'Z': 3,
}

func main() {
	guides, err := parseInput("./input-example")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	fmt.Println(calculateScorePartOne(guides))
	fmt.Println(calculateScorePartTwo(guides))
}

func parseInput(file string) ([]guide, error) {
	out := make([]guide, 0)
	fileHandle, err := os.Open(file)
	if err != nil {
		return []guide{}, err
	}

	s := bufio.NewScanner(fileHandle)
	var cur guide
	var left string
	var right string
	for s.Scan() {
		line := s.Text()
		_, err := fmt.Sscanf(line, "%v %v", &left, &right)
		if err != nil {
			log.Fatal(err)
		}
		cur.l = rune(left[0])
		cur.r = rune(right[0])
		out = append(out, cur)
	}
	return out, nil
}

func calculateScorePartOne(gs []guide) (sum int) {

	for _, g := range gs {
		if score[g.l] == score[g.r] {
			sum += 3
		}
		if score[g.r] == 1 && score[g.l] == 3 ||
			score[g.r] == 2 && score[g.l] == 1 ||
			score[g.r] == 3 && score[g.l] == 2 {
			sum += 6
		}
		sum += score[g.r]
	}
	return sum
}

func calculateScorePartTwo(gs []guide) (sum int) {

	for _, g := range gs {
		switch g.r {
		case 'X':
			// Loose
			switch score[g.l] {
			case 1:
				sum += 3
			case 2:
				sum += 1
			case 3:
				sum += 2
			}
		case 'Y':
			// Draw
			sum += 3 + score[g.l]
		case 'Z':
			// Win
			switch score[g.l] {
			case 1:
				sum += 2
			case 2:
				sum += 3
			case 3:
				sum += 1
			}
			sum += 6
		}

	}
	return sum
}
