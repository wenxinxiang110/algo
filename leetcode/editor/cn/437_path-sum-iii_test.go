package main

import (
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil root",
			args: args{
				root:      nil,
				targetSum: 8,
			},
			want: 0,
		},
		{
			name: "single node matching target",
			args: args{
				root:      &TreeNode{Val: 5},
				targetSum: 5,
			},
			want: 1,
		},
		{
			name: "single node not matching target",
			args: args{
				root:      &TreeNode{Val: 5},
				targetSum: 3,
			},
			want: 0,
		},
		{
			name: "example 1 from problem",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: -2,
							},
						},
						Right: &TreeNode{
							Val: 2,
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: -3,
						Right: &TreeNode{
							Val: 11,
						},
					},
				},
				targetSum: 8,
			},
			want: 3,
		},
		{
			name: "example 2 from problem",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: 3,
		},
		{
			name: "simple two-level tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				targetSum: 3,
			},
			want: 2, // 路径: [1,2] 和 [3]
		},
		{
			name: "negative values",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -1,
						Left: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				targetSum: 2,
			},
			want: 3, // 路径: [1,-1,2], [2], [1,-1,2] , 的另一种组合
		},
		{
			name: "complex tree with multiple paths",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				targetSum: 7,
			},
			want: 2, // 路径: [1,2,4] 和 [7]
		},
		{
			name: "zero target sum",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
				targetSum: 0,
			},
			want: 6, // 路径: [0], [0], [0], [0,0], [0,0], [0,0,0]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
