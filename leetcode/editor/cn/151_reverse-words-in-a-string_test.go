package main

import (
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "test2",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		}, {
			name: "test3",
			args: args{
				s: "F R  I   E    N     D      S      ",
			},
			want: "S D N E I R F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeExtraSpaces(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name string
		args string
		want string
	}{{
		name: "not space",
		args: "hello",
		want: "hello",
	}, {
		name: "with start space",
		args: " hello",
		want: "hello",
	}, {
		name: "mulit start space",
		args: "  hello",
		want: "hello",
	}, {
		name: "with end space",
		args: "hello ",
		want: "hello",
	}, {
		name: "with start and end space",
		args: " hello ",
		want: "hello",
	}, {
		name: "with middle space",
		args: "hello  world",
		want: "hello world",
	}, {
		name: "with middle space",
		args: "the  sky  is  blue",
		want: "the sky is blue",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeExtraSpaces([]byte(tt.args))
			if string(got) != tt.want {
				t.Errorf("reverseWords() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func Test_removeSpaces(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name string
		args string
		want string
	}{{
		name: "not space",
		args: "hello",
		want: "hello",
	}, {
		name: "with start space",
		args: " hello",
		want: "hello",
	}, {
		name: "mulit start space",
		args: "  hello",
		want: "hello",
	}, {
		name: "with end space",
		args: "hello ",
		want: "hello",
	}, {
		name: "with start and end space",
		args: " hello ",
		want: "hello",
	}, {
		name: "with middle space",
		args: "hello  world",
		want: "helloworld",
	}, {
		name: "with middle space",
		args: "the  sky  is  blue",
		want: "theskyisblue",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeSpaces([]byte(tt.args))
			if string(got) != tt.want {
				t.Errorf("removeSpaces() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
