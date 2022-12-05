package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type in struct {
	quantity int
	from     int
	to       int
}

func main() {
	rs, ins, err := parseInput("./input")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	rs2 := make([]runeStack, len(rs))
	copy(rs2, rs)
	out := applyInstructionPartOne(ins, rs)
	out.Print()
	out = applyInstructionPartTwo(ins, rs2)
	out.Print()

}

func parseInput(file string) (rs []runeStack, ins []in, err error) {
	rs = make([]runeStack, 0)
	ins = make([]in, 0)

	fileHandle, err := os.Open(file)
	if err != nil {
		return []runeStack{}, []in{}, err
	}

	s := bufio.NewScanner(fileHandle)
	var curIn in
	seenFirstLine := false
	for s.Scan() {
		line := s.Text()
		if isStackSketch(line) {
			for i := 1; i <= len(line); i += 4 {
				if i == 1 && !seenFirstLine {
					rs = make([]runeStack, (len(line)/4)+1)
					seenFirstLine = true
				}
				r := rune(line[i])
				if r != ' ' {
					rs[i/4].Push(r)
				}

			}
		}
		if len(line) > 0 && line[0] == 'm' {
			_, err := fmt.Sscanf(
				line, "move %v from %v to %v",
				&curIn.quantity, &curIn.from, &curIn.to)
			if err != nil {
				log.Fatal(err)
			}
			ins = append(ins, curIn)
		}
	}
	for _, r := range rs {
		r.Reverse()
	}
	return rs, ins, nil
}

func isStackSketch(line string) bool {
	for _, c := range line {
		if c == '[' {
			return true
		}
	}
	return false
}

func applyInstructionPartOne(ins []in, rs []runeStack) (out runeStack) {
	for _, in := range ins {
		for i := 0; i < in.quantity; i++ {
			cur := rs[in.from-1].Pop()
			rs[in.to-1].Push(cur)
		}
	}
	for _, r := range rs {
		out.Push(r.Pop())
	}
	return out
}

func applyInstructionPartTwo(ins []in, rs []runeStack) (out runeStack) {
	for _, in := range ins {
		tmp := runeStack{}
		for i := 0; i < in.quantity; i++ {
			cur := rs[in.from-1].Pop()
			tmp.Push(cur)
		}
		for !tmp.IsEmpty() {
			rs[in.to-1].Push(tmp.Pop())
		}
	}
	for _, r := range rs {
		out.Push(r.Pop())
	}
	return out
}
