package RemoveElement_27

import (
	"reflect"
	"testing"
)

func Test_RemoveElement(t *testing.T) {
	testFunctions := map[string]func(array []int, target int) int{
		"RemoveElement": RemoveElement,
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
			args: args{
				array:  []int{3, 2, 2, 3},
				target: 3,
			},
			want: 2,
		},
		{
			args: args{
				array:  []int{0, 1, 2, 2, 3, 0, 4, 2},
				target: 2,
			},
			want: 5,
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
