package PalindromeCheck

import (
	"reflect"
	"testing"
)

func Test_PalindromeCheck(t *testing.T) {
	testFunctions := map[string]func(str string) bool{
		"PalindromeCheck_V1": PalindromeCheck,
	}

	type args struct {
		str string
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				str: "abcdcba",
			},
			want: true,
		},
		{
			args: args{
				str: "abccba",
			},
			want: true,
		},
		{
			args: args{
				str: "gdacqx",
			},
			want: false,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.str); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
