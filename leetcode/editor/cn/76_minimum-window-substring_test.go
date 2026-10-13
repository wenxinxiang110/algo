package main

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			args: args{
				s: "",
				t: "",
			},
			wantAns: "",
		}, {
			name: "test1",
			args: args{
				s: "a",
				t: "a",
			},
			wantAns: "a",
		}, {
			name: "test2",
			args: args{
				s: "a",
				t: "aa",
			},
			wantAns: "",
		}, {
			name: "test3",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			wantAns: "BANC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minWindow(tt.args.s, tt.args.t); gotAns != tt.wantAns {
				t.Errorf("minWindow() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
