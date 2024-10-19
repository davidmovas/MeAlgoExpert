package twoNumberSum

import (
	algox "algoexpert/additional"
	"reflect"
	"testing"
)

func Test_twoNumberSum(t *testing.T) {

	testFunction := map[string]func(array []int, target int) []int{
		"TwoNumberSumV1": TwoNumberSumV1,
		"TwoNumberSumV2": TwoNumberSumV2,
		"TwoNumberSumV3": TwoNumberSumV3,
	}
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				array:  []int{3, 5, -4, 8, 11, 1, -1, 6},
				target: 10,
			},
			want: algox.ArraySum([]int{11, -1}),
		},
		{
			args: args{
				array:  []int{4, 6, 7, -4, 7, 1, 9},
				target: 14,
			},
			want: algox.ArraySum([]int{7, 7}),
		},
		{
			args: args{
				array:  []int{10, 5, 8, 1, -4, 7, -3, 11},
				target: 17,
			},
			want: algox.ArraySum([]int{10, 7}),
		},
	}

	for name, fn := range testFunction {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := algox.ArraySum(fn(tt.args.array, tt.args.target)); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
