package main

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1：ABCCED",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "示例2：SEE",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "SEE",
			},
			want: true,
		},
		{
			name: "示例3：ABCB",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			want: false,
		},
		{
			name: "空网格",
			args: args{
				board: [][]byte{},
				word:  "TEST",
			},
			want: false,
		},
		{
			name: "空单词",
			args: args{
				board: [][]byte{
					{'A', 'B'},
					{'C', 'D'},
				},
				word: "",
			},
			want: false,
		},
		{
			name: "单字符网格",
			args: args{
				board: [][]byte{
					{'X'},
				},
				word: "X",
			},
			want: true,
		},
		{
			name: "复杂路径测试",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C'},
					{'D', 'E', 'F'},
					{'G', 'H', 'I'},
				},
				word: "ADGHI",
			},
			want: true,
		},
		{
			name: "重复字符测试",
			args: args{
				board: [][]byte{
					{'A', 'A', 'A'},
					{'A', 'A', 'A'},
					{'A', 'A', 'A'},
				},
				word: "AAAAAAAAA",
			},
			want: true,
		},
		{
			name: "不可达路径",
			args: args{
				board: [][]byte{
					{'A', 'B'},
					{'C', 'D'},
				},
				word: "ACB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
