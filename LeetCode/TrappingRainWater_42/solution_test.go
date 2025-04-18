package TrappingRainWater_42

import (
	"reflect"
	"testing"
)

func Test_TrappingRainWater(t *testing.T) {
	testFunctions := map[string]func([]int) int{
		"TrappingRainWater_V1": v1,
	}

	type args struct {
		height []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.height); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
