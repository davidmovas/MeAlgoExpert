package ValidateSubsequence

import (
	"reflect"
	"testing"
)

func Test_validateSubsequence(t *testing.T) {
	testFunction := map[string]func(array []int, sequence []int) bool{
		"ValidateSubsequenceV1": ValidateSubsequenceV1,
	}

	type args struct {
		array    []int
		sequence []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
				sequence: []int{1, 6, -1, 10},
			},
			want: true,
		},
		{
			args: args{
				array:    []int{12, 5, 22, 25, 6, 9, 8, 10},
				sequence: []int{1, 6, -1, 10},
			},
			want: false,
		},
	}

	for name, fn := range testFunction {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.array, tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
