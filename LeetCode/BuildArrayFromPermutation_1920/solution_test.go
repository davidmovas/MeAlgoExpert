package BuildArrayFromPermutation_1920

import (
	"reflect"
	"testing"
)

func Test_BuildArrayFromPermutation(t *testing.T) {
	testFunctions := map[string]func(nums []int) []int{
		"BuildArrayFromPermutation_V1": v1,
	}

	type args struct {
		nums []int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{0, 2, 1, 5, 3, 4},
			},
			want: []int{0, 1, 2, 4, 5, 3},
		},
		{
			args: args{
				nums: []int{5, 0, 1, 2, 3, 4},
			},
			want: []int{4, 5, 0, 1, 2, 3},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
