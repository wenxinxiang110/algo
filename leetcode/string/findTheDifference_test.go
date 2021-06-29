package string

import (
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"",
			args{
				s: "abcd",
				t: "abcde",
			},
			'e',
		},
		{
			"",
			args{
				s: "",
				t: "y",
			},
			'y',
		},
		{
			"",
			args{
				s: "a",
				t: "aa",
			},
			'a',
		},
		{
			"",
			args{
				s: "ae",
				t: "aea",
			},
			'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
