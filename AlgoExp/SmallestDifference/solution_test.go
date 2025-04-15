package SmallestDifference

import (
	"reflect"
	"testing"
)

func Test_SmallestDifference(t *testing.T) {
	testFunctions := map[string]func(first, second []int) []int{
		"SmallestDifference_V1": SmallestDifference,
	}

	type args struct {
		first  []int
		second []int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				first:  []int{-1, 5, 10, 20, 28, 3},
				second: []int{26, 134, 135, 15, 17},
			},
			want: []int{28, 26},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
