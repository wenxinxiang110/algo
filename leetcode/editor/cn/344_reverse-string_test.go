package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {

	tests := []struct {
		name string
		args []byte
		want []byte
	}{
		{
			name: "test1",
			args: []byte("hello"),
			want: []byte("olleh"),
		}, {
			name: "test2",
			args: []byte("world"),
			want: []byte("dlrow"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
