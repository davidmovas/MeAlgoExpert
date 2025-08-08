package ConcatenationOfArray_1929

import (
	"reflect"
	"testing"
)

func Test_ConcatenationOfArray(t *testing.T) {
	testFunctions := map[string]func(naum []int) []int{
		"ConcatenationOfArray_V1": v1,
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
				nums: []int{1, 2, 1},
			},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			args: args{
				nums: []int{1, 3, 2, 1},
			},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
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
