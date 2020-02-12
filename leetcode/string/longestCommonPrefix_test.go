package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case-1",
			args{[]string{"flower", "flow", "flight"}},
			"fl",
		},
		{
			"case-2",
			args{[]string{"dog", "racecar", "car"}},
			"",
		},
		{
			"case-3",
			args{[]string{"ca", "a"}},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindCom(t *testing.T) {
	assert.Equal(t, "", FindCom("qwe", "asd"))

	assert.Equal(t, "asd", FindCom("asdert", "asd"))

	assert.Equal(t, "asd", FindCom("asdqwe", "asdzxc"))

	assert.Equal(t, "", FindCom("ca", "a"))

}

func TestLCP(t *testing.T) {
	type args struct {
		this   string
		others []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			"case-1",
			args{
				this:   "flower",
				others: []string{"flow", "flight"},
			},
			"fl",
		},
		{
			"",
			args{
				this:   "dog",
				others: []string{"racecar", "car"},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := LCP(tt.args.this, tt.args.others); gotResult != tt.wantResult {
				t.Errorf("LCP() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
