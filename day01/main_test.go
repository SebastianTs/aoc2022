package main

import "testing"

func TestParseInputAtoi(t *testing.T) {

	_, err := parseInputAtoi("input")
	if err != nil {
		t.Errorf("%s", err)
	}
}

func BenchmarkTestParseInputAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInputAtoi("input")
	}
}

func TestParseInputScan(t *testing.T) {
	data := []struct {
		inputFile string
	}{
		{"./input"},
	}
	for _, d := range data {
		_, err := parseInputScan(d.inputFile)
		if err != nil {
			t.Errorf("%s", err)
		}
	}
}

func BenchmarkTestParseInputScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInputScan("input")
	}
}
