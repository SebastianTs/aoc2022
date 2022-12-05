package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type instruction struct {
	b1 int
	e1 int
	b2 int
	e2 int
}

func main() {
	ins, err := parseInput("./input")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	fmt.Println(countCover(ins, isCompleteCover))
	fmt.Println(countCover(ins, hasPartialCover))
}

func parseInput(file string) ([]instruction, error) {
	out := make([]instruction, 0)
	fileHandle, err := os.Open(file)
	if err != nil {
		return []instruction{}, err
	}

	s := bufio.NewScanner(fileHandle)
	var cur instruction
	var b1, b2 int
	var e1, e2 int
	for s.Scan() {
		line := s.Text()
		_, err := fmt.Sscanf(line, "%v-%v,%v-%v", &b1, &e1, &b2, &e2)
		if err != nil {
			log.Fatal(err)
		}
		cur.b1 = b1
		cur.e1 = e1
		cur.b2 = b2
		cur.e2 = e2
		out = append(out, cur)
	}
	return out, nil
}

func countCover(is []instruction, f1 func(i instruction) bool) (sum int) {
	for _, in := range is {
		if f1(in) {
			sum++
		}
	}
	return sum
}

func isCompleteCover(i instruction) bool {
	return (i.e2 <= i.e1 && i.b2 >= i.b1) || (i.e1 <= i.e2 && i.b1 >= i.b2)
}

func hasPartialCover(i instruction) bool {
	return !hasNoCoverage(i)
}

func hasNoCoverage(i instruction) bool {
	return i.b2 > i.e1 || i.b1 > i.e2
}
