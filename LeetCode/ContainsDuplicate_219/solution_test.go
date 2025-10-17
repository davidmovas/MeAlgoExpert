package ContainsDuplicate_219

import (
	"reflect"
	"testing"
)

func Test_ContainsDuplicate(t *testing.T) {
	testFunctions := map[string]func(nums []int, k int) bool{
		"ContainsDuplicate_V1": v1,
	}

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
