package ContainerWithMostWater_11

import (
	"reflect"
	"testing"
)

func Test_ContainerWithMostWater(t *testing.T) {
	testFunctions := map[string]func(height []int) int{
		"ContainerWithMostWater_V1": v1,
	}

	type args struct {
		height []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.height); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
