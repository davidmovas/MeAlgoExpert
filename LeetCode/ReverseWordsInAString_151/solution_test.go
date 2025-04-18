package ReverseWordsInAString_151

import (
	"reflect"
	"testing"
)

func Test_ReverseWordsInAString(t *testing.T) {
	testFunctions := map[string]func(string) string{
		"ReverseWordsInAString_V1": v1,
	}

	type args struct {
		input string
	}

	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				input: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			args: args{
				input: "  hello world  ",
			},
			want: "world hello",
		},
		{
			args: args{
				input: "a good   example",
			},
			want: "example good a",
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
