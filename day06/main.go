package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	ins, err := parseInput("./input")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	for _, s := range ins {
		fmt.Println(calculateMarker(s, 4))
		fmt.Println(calculateMarker(s, 14))
	}
}

func parseInput(file string) (ins []string, err error) {
	ins = make([]string, 0)

	fileHandle, err := os.Open(file)
	if err != nil {
		return []string{}, err
	}

	s := bufio.NewScanner(fileHandle)
	for s.Scan() {
		ins = append(ins, s.Text())
	}
	return ins, nil
}

func calculateMarker(s string, ws int) int {

	for i := ws; i < len(s); i++ {
		set := make(map[byte]struct{})
		for j := 0; j < ws; j++ {
			set[s[i-j]] = struct{}{}
		}
		if len(set) == ws {
			return i + 1
		}
	}
	return 0
}
