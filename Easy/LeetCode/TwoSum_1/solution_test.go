package TwoSum_1

import (
	"reflect"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	testFunctions := map[string]func(nums []int, target int) []int{
		"TwoSum_V1": v1,
		//"TwoSum_V2": v2,
		"TwoSum_V2": v3,
	}

	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
