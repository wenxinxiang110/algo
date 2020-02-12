package string

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name    string
		args    args
		wantSum string
	}{
		{
			"case-1",
			args{
				a: "11",
				b: "1",
			},
			"100",
		},
		{
			"case-2",
			args{
				a: "1010",
				b: "1011",
			},
			"10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := addBinary(tt.args.a, tt.args.b); gotSum != tt.wantSum {
				t.Errorf("addBinary() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
