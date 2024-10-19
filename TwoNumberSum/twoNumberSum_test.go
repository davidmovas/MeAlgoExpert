package main

import (
	"reflect"
	"testing"
)

func Test_twoNumberSum(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				array:  []int{3, 5, -4, 8, 11, 1, -1, 6},
				target: 10,
			},
			want: []int{11, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoNumberSum(tt.args.array, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoNumberSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
