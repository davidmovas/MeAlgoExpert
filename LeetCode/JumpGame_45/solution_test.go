package JumpGame_55

import (
	"reflect"
	"testing"
)

func Test_JumpGame(t *testing.T) {
	testFunctions := map[string]func(nums []int) int{
		"JumpGame_V1": v1,
	}

	type args struct {
		nums []int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
