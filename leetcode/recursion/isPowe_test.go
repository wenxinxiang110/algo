package recursion

import (
	"testing"
)

func Test_isPowerOfTwo(t *testing.T) {
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

func Test_isPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		//{
		//	"",
		//	1,
		//	true,
		//},
		{
			"",
			27,
			true,
		},
		//{
		//	"",
		//	0,
		//	false,
		//},
		//{
		//	"",
		//	9,
		//	true,
		//},
		//{
		//	"",
		//	45,
		//	false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args); got != tt.want {
				t.Errorf("isPowerOfThree(%d) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
