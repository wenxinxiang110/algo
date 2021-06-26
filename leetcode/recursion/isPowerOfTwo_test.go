package recursion

import (
	"testing"
)

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			"",
			1,
			true,
		},
		{
			"",
			16,
			true,
		},
		{
			"",
			3,
			false,
		},
		{
			"",
			4,
			true,
		},
		{
			"",
			5,
			false,
		},
		{
			"",
			-16,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args); got != tt.want {
				t.Errorf("isPowerOfTwo(%d) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
