package GasStation_134

import (
	"reflect"
	"testing"
)

func Test_GasStation(t *testing.T) {
	testFunctions := map[string]func([]int, []int) int{
		"GasStation_V1": v1,
	}

	type args struct {
		gas, cost []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			want: -1,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.gas, tt.args.cost); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
