package HIndex_274

import (
	"reflect"
	"testing"
)

func Test_HIndex(t *testing.T) {
	testFunctions := map[string]func(citations []int) int{
		"HIndex_V1": v1,
	}

	type args struct {
		citations []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				citations: []int{3, 0, 6, 1, 5},
			},
			want: 3,
		},
		{
			args: args{
				citations: []int{1, 3, 1},
			},
			want: 1,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.citations); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
