package dp

import (
	"testing"
)

func TestBasePackage(t *testing.T) {
	type args struct {
		items     []Item
		maxWeight int
	}
	tests := []struct {
		name         string
		args         args
		wantMaxValue int
	}{
		{
			name: "base package",
			args: args{
				items: []Item{
					{Weight: 1, Value: 15},
					{Weight: 3, Value: 20},
					{Weight: 4, Value: 30},
				},
				maxWeight: 4,
			},
			wantMaxValue: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxValue := BasePackage(tt.args.items, tt.args.maxWeight); gotMaxValue != tt.wantMaxValue {
				t.Errorf("BasePackage() = %v, want %v", gotMaxValue, tt.wantMaxValue)
			}
		})
	}
}
