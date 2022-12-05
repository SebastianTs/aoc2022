package main

import "testing"

func Test_calculateScorePartOne(t *testing.T) {
	type args struct {
		gs []guide
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{"Example", args{[]guide{{'A', 'Y'}, {'B', 'X'}, {'C', 'Z'}}}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := calculateScorePartOne(tt.args.gs); gotSum != tt.wantSum {
				t.Errorf("calculateScorePartOne() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
