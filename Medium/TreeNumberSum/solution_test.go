package TreeNumberSum

import (
	"reflect"
	"testing"
)

func Test_TreeNumberSum(t *testing.T) {
	testFunctions := map[string]func(array []int, target int) [][]int{
		"TreeNumberSum_V1": TreeNumberSum,
	}

	type args struct {
		array  []int
		target int
	}

	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{
				array:  []int{12, 3, 1, 2, -6, 5, 8, 6},
				target: 0,
			},
			want: [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}},
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
