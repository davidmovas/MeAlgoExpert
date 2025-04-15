package PascalsTriangle_118

import (
	"reflect"
	"testing"
)

func Test_PascalTriangle(t *testing.T) {
	testFunctions := map[string]func(numRows int) [][]int{
		"PascalTriangle_V1": v1,
	}

	type args struct {
		numRows int
	}

	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{
				numRows: 5,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			args: args{
				numRows: 1,
			},
			want: [][]int{{1}},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
