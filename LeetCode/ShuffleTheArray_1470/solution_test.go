package ShuffleTheArray_1470

import (
	"reflect"
	"testing"
)

func Test_ShuffleTheArray(t *testing.T) {
	testFunctions := map[string]func(nums []int, n int) []int{
		"ShuffleTheArray_V1": v1,
	}

	type args struct {
		nums []int
		n    int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				[]int{2, 5, 1, 3, 4, 7},
				3,
			},
			want: []int{2, 3, 5, 4, 1, 7},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
