package JumpGame_55

import (
	"reflect"
	"testing"
)

func Test_JumpGame(t *testing.T) {
	testFunctions := map[string]func(nums []int) bool{
		"JumpGame_V1": v1,
	}

	type args struct {
		nums []int
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
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
