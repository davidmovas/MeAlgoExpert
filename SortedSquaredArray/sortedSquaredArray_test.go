package SortedSquaredArray

import (
	"reflect"
	"testing"
)

func Test_sortedSquaredArray(t *testing.T) {
	testFunctions := map[string]func(array []int) []int{
		"SortedSquaredArray_V1": SortedSquaredArrayV1,
	}

	type args struct {
		array []int
	}
	tests := []struct {
		args   args
		want   []int
		result bool
	}{
		{
			args: args{
				array: []int{1, 2, 3, 5, 6, 8, 9},
			},
			want:   []int{1, 4, 9, 25, 36, 64, 81},
			result: true,
		},
		{
			args: args{
				array: []int{6, 2, 9, 4, 1, 5, 100},
			},
			want:   []int{1, 4, 9, 25, 36, 64, 81},
			result: false,
		},
		{
			args: args{
				array: []int{4, 3, 8, 2},
			},
			want:   []int{16, 9, 64, 4},
			result: true,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.array); !reflect.DeepEqual(got, tt.want) {
					if tt.result {
						t.Errorf("%s = %v, want %v", name, got, tt.want)
					}
				}
			}
		})
	}
}
