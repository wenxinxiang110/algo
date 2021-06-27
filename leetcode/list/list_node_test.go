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
			"nil",
			args{
				exp: "",
				spt: "->",
			},
			nil,
		},
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

func TestListNode_String(t *testing.T) {

	tests := []struct {
		name    string
		linkExp string
	}{
		// TODO: Add test cases.
		{
			"nil",
			"",
		},
		{
			"normal",
			"1->1->2",
		},
		{
			"normal",
			"1->1->2->3->3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleParse(tt.linkExp).String(); got != tt.linkExp {
				t.Errorf("String(%v) = %v", tt.linkExp, got)
			}
		})
	}
}
