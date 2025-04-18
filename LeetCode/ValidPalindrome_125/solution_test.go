package ValidPalindrome_125

import (
	"reflect"
	"testing"
)

func Test_ValidPalindrome(t *testing.T) {
	testFunctions := map[string]func(string) bool{
		"ValidPalindrome_V1": v1,
	}

	type args struct {
		input string
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				input: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			args: args{
				input: "race a car",
			},
			want: false,
		},
		{
			args: args{
				input: " ",
			},
			want: true,
		},
		{
			args: args{
				input: "0P",
			},
			want: false,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.input); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
