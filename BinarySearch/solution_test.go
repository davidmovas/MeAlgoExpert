package BinarySearch

import (
	"reflect"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	testFunctions := map[string]func(array []int, target int) int{
		"BinarySearch_V1": BinarySearch,
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
			args: args{array: []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, target: 33},
			want: 3,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
