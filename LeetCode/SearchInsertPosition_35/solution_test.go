package SearchInsertPosition_35

import (
	"reflect"
	"testing"
)

func Test_SearchInsertPosition(t *testing.T) {
	testFunctions := map[string]func(nums []int, target int) int{
		"SearchInsertPosition_V1": v1,
	}

	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
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
