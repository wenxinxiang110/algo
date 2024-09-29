package main

import (
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		}, {
			name: "test 2",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "bacd",
		}, {
			name: "test 3",
			args: args{
				s: "a",
				k: 2,
			},
			want: "a",
		}, {
			name: "test 4",
			args: args{
				s: "abcdefg",
				k: 8,
			},
			want: "gfedcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStrV2(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
