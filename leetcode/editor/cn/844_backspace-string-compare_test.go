package main

import (
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			//	测试用例:"bxj##tw"
			//			"bxj###tw"
			name: "test5",
			args: args{
				s: "bxj##tw",
				t: "bxj###tw",
			},
			want: false,
		},
		{
			//	测试用例:"a"
			//			"aa#a"
			name: "test6",
			args: args{
				s: "a",
				t: "aa#a",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
