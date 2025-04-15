package RotateArray_189

import (
	"reflect"
	"testing"
)

func Test_RotateArray(t *testing.T) {
	testFunctions := map[string]func(array []int, k int) []int{
		"RotateArray_V1": v1,
		//"RotateArray_V2": v2,
		"RotateArray_V3": v3,
	}

	type args struct {
		array []int
		k     int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7},
				k:     3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			args: args{
				array: []int{-1, -100, 3, 99},
				k:     2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			args: args{
				array: []int{1, 2},
				k:     3,
			},
			want: []int{2, 1},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.array, tt.args.k); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
