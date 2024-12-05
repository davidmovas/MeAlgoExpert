package ProductSum

import (
	"reflect"
	"testing"
)

func Test_ProductSum(t *testing.T) {
	testFunctions := map[string]func(array []any) int{
		"ProductSum_V1": ProductSum,
	}

	type args struct {
		array []any
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{[]any{5, 2, []any{7, -1}, 3, []any{6, []any{-13, 8}, 4}}},
			want: 12,
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
