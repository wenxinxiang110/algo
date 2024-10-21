package list

import (
	"testing"

	"github.com/NothingXiang/algo/leetcode/list_node"
)

func Test_reverseList(t *testing.T) {

	tests := []struct {
		name string
		args *ListNode
		want *ListNode
	}{
		{
			"nil",
			SimpleParse(""),
			SimpleParse(""),
		},
		{
			"normal",
			SimpleParse("1->2"),
			SimpleParse("2->1"),
		},
		{
			"normal",
			SimpleParse("1->2->3->4->5"),
			SimpleParse("5->4->3->2->1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args); !list_node.Equal(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
