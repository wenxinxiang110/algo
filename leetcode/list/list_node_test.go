package list

import (
	"reflect"
	"testing"
)

func TestParseListNode(t *testing.T) {
	type args struct {
		exp string
		spt string
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"",
			args{
				exp: "1->1->2",
				spt: "->",
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 2},
				},
			},
		},
		{
			"",
			args{
				exp: "1->1->2->3->3",
				spt: "->",
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseListNode(tt.args.exp, tt.args.spt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
