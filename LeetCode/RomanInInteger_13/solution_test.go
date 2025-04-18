package RomanInInteger_13

import (
	"reflect"
	"testing"
)

func Test_RomanToInteger(t *testing.T) {
	testFunctions := map[string]func(string) int{
		"RomanToInteger_V1": v1,
	}

	type args struct {
		roman string
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				roman: "III",
			},
			want: 3,
		},
		{
			args: args{
				roman: "LVIII",
			},
			want: 58,
		},
		{
			args: args{
				roman: "MCMXCIV",
			},
			want: 1994,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.roman); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
