package TransposeMatrix

import (
	"reflect"
	"testing"
)

func Test_transposeMatrix(t *testing.T) {
	testFunctions := map[string]func(matrix [][]int) [][]int{
		"transposeMatrix_V1": TransposeMatrix,
	}

	type args struct {
		matrix [][]int
	}

	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{
				matrix: [][]int{
					{1, 2},
				},
			},
			want: [][]int{{1}, {2}},
		},
		{
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
			},
			want: [][]int{
				{1, 3, 5},
				{2, 4, 6},
			},
		},
		{
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
