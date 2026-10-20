package tree

import (
	"testing"
)

func TestTreeNode_PreorderTraversal(t *testing.T) {

	var data = make(map[string][]int)
	visitor := func(node *TreeNode) {

	}
	type args struct {
		visit Visitor
	}
	tests := []struct {
		name   string
		fields *TreeNode
		args   args
	}{
		{
			name: "nil", fields: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.fields
			root.PreorderTraversal(tt.args.visit)
		})
	}
}
