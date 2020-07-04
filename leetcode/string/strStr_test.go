package string

import (
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				haystack: "hello",
				needle:   "ll",
			},
			2,
		},
		{
			"",
			args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			-1,
		},
		{
			"",
			args{
				haystack: "",
				needle:   "",
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
