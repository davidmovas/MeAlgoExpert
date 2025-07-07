package TwoSum2_167

import (
	"reflect"
	"testing"
)

func Test_TwoPointersSum2(t *testing.T) {
	testFunctions := map[string]func(numbers []int, target int) []int{
		"TwoPointersSum2_V1": v1,
	}

	type args struct {
		numbers []int
		target  int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
