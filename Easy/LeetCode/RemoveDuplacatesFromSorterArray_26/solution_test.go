package RemoveDuplacatesFromSorterArray_26

import (
	"reflect"
	"testing"
)

func Test_RemoveDuplicatesFromSortedArray(t *testing.T) {
	testFunctions := map[string]func(array []int) int{
		"RemoveDuplicatesFromSortedArray": RemoveDuplicatesFromSortedArray,
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
				array: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			args: args{
				array: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
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
