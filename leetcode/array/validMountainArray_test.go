package array

import (
	"testing"
)

func Test_validMountainArray(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want bool
	}{
		// TODO: Add test cases.
		{
			"输入：arr = [2,1]\n输出：false",
			[]int{2, 1},
			false,
		},
		{
			"输入：arr = [3,5,5]\n输出：false",
			[]int{3, 5, 5},
			false,
		},
		{
			"输入：arr = [0,3,2,1]\n输出：true",
			[]int{0, 3, 2, 1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
