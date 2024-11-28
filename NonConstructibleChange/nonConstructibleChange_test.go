package NonConstructibleChange

import (
	"reflect"
	"testing"
)

func Test_nonConstructibleChange(t *testing.T) {
	testFunctions := map[string]func(coins []int) int{
		"nonConstructibleChange_V1": NonConstructibleChange,
	}

	type args struct {
		array []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				array: []int{5, 7, 1, 1, 2, 3, 22},
			},
			want: 20,
		},
		{
			args: args{
				array: []int{1, 1, 1, 1, 1},
			},
			want: 6,
		},
	}
	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.array); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
