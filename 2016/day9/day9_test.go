package main

import (
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		sequence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Sequence 1", args{"ADVENT"}, "ADVENT"},
		{"Sequence 2", args{"A(1x5)BC"}, "ABBBBBC"},
		{"Sequence 3", args{"(3x3)XYZ"}, "XYZXYZXYZ"},
		{"Sequence 4", args{"A(2x2)BCD(2x2)EFG"}, "ABCBCDEFEFG"},
		{"Sequence 5", args{"(6x1)(1x3)A"}, "(1x3)A"},
		{"Sequence 6", args{"X(8x2)(3x3)ABCY"}, "X(3x3)ABC(3x3)ABCY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.sequence); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeLengthVersion2(t *testing.T) {
	type args struct {
		sequence string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sequence 1", args{"(3x3)XYZ"}, 9},
		{"Sequence 2", args{"X(8x2)(3x3)ABCY"}, 20},
		{"Sequence 3", args{"(27x12)(20x12)(13x14)(7x10)(1x12)A"}, 241920},
		{"Sequence 4", args{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"}, 445},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeLengthVersion2(tt.args.sequence); got != tt.want {
				t.Errorf("decodeLengthVersion2() = %v, want %v", got, tt.want)
			}
		})
	}
}
