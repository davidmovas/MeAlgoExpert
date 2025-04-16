package ProductOfArrayExceptSelf_238

import (
	"reflect"
	"testing"
)

func Test_ProductOfArrayExceptSelf(t *testing.T) {
	testFunctions := map[string]func([]int) []int{
		"ProductOfArrayExceptSelf_V1": v1,
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
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
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
