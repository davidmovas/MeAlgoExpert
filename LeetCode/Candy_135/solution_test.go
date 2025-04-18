package Candy_135

import (
	"reflect"
	"testing"
)

func Test_Candy(t *testing.T) {
	testFunctions := map[string]func([]int) int{
		"Candy_V1": v1,
	}

	type args struct {
		ratings []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.ratings); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
