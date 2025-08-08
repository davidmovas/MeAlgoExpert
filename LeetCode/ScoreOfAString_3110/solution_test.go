package ScoreOfAString_3110

import (
	"reflect"
	"testing"
)

func Test_ScoreOfAString(t *testing.T) {
	testFunctions := map[string]func(s string) int{
		"ScoreOfAString_V1": v1,
	}

	type args struct {
		s string
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{s: "hello"}, want: 13},
		{args: args{s: "zaz"}, want: 50},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
