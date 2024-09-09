package main

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantArray []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			wantArray: []int{0, 1, 3, 0, 4},
			want:      5,
		},
		{
			name: "case2",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			wantArray: []int{2, 2},
			want:      2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
				return
			}
			if !assertElementEqual(tt.args.nums, tt.wantArray, tt.want) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums, tt.wantArray)
			}
		})
	}
}

func assertElementEqual(got, want []int, num int) bool {
	for i := 0; i < num; i++ {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
