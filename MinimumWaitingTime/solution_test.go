package MinimumWaitingTime

import (
	"reflect"
	"testing"
)

func Test_MinimumWaitingTime(t *testing.T) {
	testFunctions := map[string]func(queries []int) int{
		"MinimumWaitingTime_V1": MinimumWaitingTime_V1,
	}

	type args struct {
		queries []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				queries: []int{3, 2, 1, 2, 6},
			},
			want: 17,
		},
		{
			args: args{
				queries: []int{1, 6},
			},
			want: 1,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
