package testutil

import (
	"testing"
)

func TestSliceEqual(t *testing.T) {
	type args[T comparable] struct {
		s1 []T
		s2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty",
			want: true,
		}, {
			name: "equal",
			args: args[int]{
				s1: []int{1, 2, 3},
				s2: []int{1, 2, 3},
			},
			want: true,
		}, {
			name: "disorder",
			args: args[int]{
				s1: []int{1, 2, 3},
				s2: []int{1, 3, 2},
			},
			want: true,
		}, {
			name: "unequal",
			args: args[int]{
				s1: []int{1, 2, 3},
				s2: []int{1, 2, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("SliceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
