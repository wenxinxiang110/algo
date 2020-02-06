package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRtoI(t *testing.T) {
	assert.Equal(t,1,RtoI('I'))

	assert.Equal(t,5,RtoI('V'))

	assert.Equal(t,10,RtoI('X'))

	assert.Equal(t,50,RtoI('L'))

	assert.Equal(t,100,RtoI('C'))

	assert.Equal(t,500,RtoI('D'))

	assert.Equal(t,1000,RtoI('M'))

}

func TestSpecialRule(t *testing.T) {

	assert.Equal(t,true,SpecialRule('I','V'))
	assert.Equal(t,true,SpecialRule('I','X'))

	assert.Equal(t,false,SpecialRule('I','O'))

}

func Test_romanToInt(t *testing.T) {

	tests := []struct {
		name       string
		args       string
		wantResult int
	}{
		// TODO: Add test cases.
		{
			"result-2",
			"II",
			2,
		},
		{
			"result-4",
			"IV",
			4,
		},
		{
			"result-2",
			"IX",
			9,
		},
		{
			"result-2",
			"II",
			2,
		},
		{
			"result-58",
			"LVIII",
			58,
		},
		{
			"result-1994",
			"MCMXCIV",
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := romanToInt(tt.args); gotResult != tt.wantResult {
				t.Errorf("romanToInt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}