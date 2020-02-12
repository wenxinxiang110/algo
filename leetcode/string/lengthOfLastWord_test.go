package string

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	/*	type args struct {
		s string
	}*/
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			"",
			"Hello World",
			5,
		},
		{
			"",
			"a ",
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
