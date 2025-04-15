package PlusOne_66

import (
	"reflect"
	"testing"
)

func Test_PlusOne(t *testing.T) {
	testFunctions := map[string]func(digits []int) []int{
		//"PlusOne_V1": v1,
		"PlusOne_V2": v2,
		"PlusOne_V3": v3,
	}

	type args struct {
		digits []int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
		{
			args: args{
				digits: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			},
			want: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
