package main

import (
	"testing"
)

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1: 可以完成",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: true,
		},
		{
			name: "示例2: 有环，不能完成",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
		{
			name: "没有先修课程",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{},
			},
			want: true,
		},
		{
			name: "单门课程",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: true,
		},
		{
			name: "线性依赖（无环）",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0},
					{2, 1},
					{3, 2},
				},
			},
			want: true,
		},
		{
			name: "复杂有环情况",
			args: args{
				numCourses: 5,
				prerequisites: [][]int{
					{1, 0},
					{2, 1},
					{3, 2},
					{1, 3}, // 形成环：0->1->2->3->1
				},
			},
			want: false,
		},
		{
			name: "多个独立课程组",
			args: args{
				numCourses: 6,
				prerequisites: [][]int{
					{1, 0},
					{2, 1},
					{4, 3},
					{5, 4},
				},
			},
			want: true,
		},
		{
			name: "自环（课程依赖自己）",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{0, 0}, // 课程0依赖自己
					{1, 0},
				},
			},
			want: false,
		},
		{
			name: "多入度节点",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{3, 0},
					{3, 1},
					{3, 2},
				},
			},
			want: true,
		},
		{
			name: "多出度节点",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0},
					{2, 0},
					{3, 0},
				},
			},
			want: true,
		},
		{
			name: "最大边界测试",
			args: args{
				numCourses:    2000,
				prerequisites: [][]int{},
			},
			want: true,
		},
		{
			name: "复杂无环图",
			args: args{
				numCourses: 7,
				prerequisites: [][]int{
					{1, 0},
					{2, 0},
					{3, 1},
					{4, 2},
					{5, 3},
					{6, 4},
					{6, 5},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
