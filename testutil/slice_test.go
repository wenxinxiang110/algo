package testutil

import (
	"fmt"
	"testing"
	"time"
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

func TestTime(t *testing.T) {
	timeStr := "2024-11-12T16:47:20+0800"
	// 定义时间格式，Go 中的时间格式使用一个固定的参考时间 "2006-01-02 15:04:05 MST"
	layout := time.RFC3339
	layoutNano := time.RFC3339Nano
	layoutCustom := "2006-01-02T15:04:05-0700"

	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	//fmt.Println("Parsed time:", parsedTime)

	parsedTime, err = time.Parse(layoutNano, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	parsedTime, err = time.Parse(layoutCustom, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println("Parsed time:", parsedTime)

}
