package main

import "fmt"

type runeStack struct {
	stackSlice []rune
}

func (s *runeStack) Push(in rune) {
	s.stackSlice = append(s.stackSlice, in)
}

func (s *runeStack) Pop() rune {
	if s.IsEmpty() {
		var r rune
		return r
	}
	out := s.stackSlice[len(s.stackSlice)-1]
	s.stackSlice = s.stackSlice[:len(s.stackSlice)-1]
	return out
}

func (s *runeStack) Reverse() {
	for i, j := 0, len(s.stackSlice)-1; i < j; i, j = i+1, j-1 {
		s.stackSlice[i], s.stackSlice[j] = s.stackSlice[j], s.stackSlice[i]
	}
}

func (s *runeStack) IsEmpty() bool {
	return len(s.stackSlice) == 0
}

func (s *runeStack) Print() {
	for _, r := range s.stackSlice {
		fmt.Printf("%v", string(r))
	}
	fmt.Print("\n")
}
