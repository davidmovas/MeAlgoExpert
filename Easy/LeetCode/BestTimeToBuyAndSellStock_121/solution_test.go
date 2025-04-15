package BestTimeToBuyAndSellStock_121

import (
	"reflect"
	"testing"
)

func Test_BestTimeToBuyAndSellStock(t *testing.T) {
	testFunctions := map[string]func(prices []int) int{
		"BestTimeToBuyAndSellStock_V1": v1,
		"BestTimeToBuyAndSellStock_V2": v2,
	}

	type args struct {
		prices []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
