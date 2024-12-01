package TandemBicycle

import (
	"reflect"
	"testing"
)

func Test_TandemBicycle(t *testing.T) {
	testFunctions := map[string]func(red, blue []int, fastest bool) int{
		"TandemBicycle_V1": TendemBicycle_V1,
	}

	type args struct {
		red, blue []int
		fastest   bool
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				red:     []int{5, 5, 3, 9, 2},
				blue:    []int{3, 6, 7, 2, 1},
				fastest: true,
			},

			want: 32,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.red, tt.args.blue, tt.args.fastest); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
