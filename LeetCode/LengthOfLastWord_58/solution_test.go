package LengthOfLastWord_58

import (
	"reflect"
	"testing"
)

func Test_LengthOfLastWord(t *testing.T) {
	testFunctions := map[string]func(string) int{
		"LengthOfLastWord_V1": v1,
		"LengthOfLastWord_V2": v2,
	}

	type args struct {
		sentence string
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				sentence: "Hello World",
			},
			want: 5,
		},
		{
			args: args{
				sentence: "   fly me   to   the moon  ",
			},
			want: 4,
		},
		{
			args: args{
				sentence: "luffy is still joyboy",
			},
			want: 6,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.sentence); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
