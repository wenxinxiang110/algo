package main

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1: ababcbacadefegdehijhklij",
			args: args{s: "ababcbacadefegdehijhklij"},
			want: []int{9, 7, 8},
		},
		{
			name: "示例2: eccbbbbdec",
			args: args{s: "eccbbbbdec"},
			want: []int{10},
		},
		{
			name: "单字符字符串",
			args: args{s: "a"},
			want: []int{1},
		},
		{
			name: "全相同字符",
			args: args{s: "aaaaa"},
			want: []int{5},
		},
		{
			name: "交替字符ababab",
			args: args{s: "ababab"},
			want: []int{6},
		},
		{
			name: "三个独立片段",
			args: args{s: "abc"},
			want: []int{1, 1, 1},
		},
		{
			name: "重叠字符模式",
			args: args{s: "abacaba"},
			want: []int{7},
		},
		{
			name: "复杂重叠模式",
			args: args{s: "ababcbacadefegde"},
			want: []int{9, 7},
		},
		{
			name: "多个独立片段",
			args: args{s: "abcdefghij"},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "嵌套重叠模式",
			args: args{s: "qiejxqfnqceocmy"},
			want: []int{13, 1, 1},
		},
		{
			name: "边界情况-最小长度",
			args: args{s: "a"},
			want: []int{1},
		},
		{
			name: "重复字符组合",
			args: args{s: "aabbccdd"},
			want: []int{2, 2, 2, 2},
		},
		{
			name: "复杂字母分布",
			args: args{s: "caedbdedda"},
			want: []int{1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
