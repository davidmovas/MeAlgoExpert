package FindTheIndexFirstOccurrenceInAString_28

import (
	"reflect"
	"testing"
)

func Test_FindIndexOfSubstring(t *testing.T) {
	testFunctions := map[string]func(string, string) int{
		"FindIndexOfSubstring_V1": v1,
		"FindIndexOfSubstring_V2": v2,
	}

	type args struct {
		str, sub string
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				str: "sadbutsad",
				sub: "sad",
			},
			want: 0,
		},
		{
			args: args{
				str: "leetcode",
				sub: "leeto",
			},
			want: -1,
		},
		{
			args: args{
				str: "mississippi",
				sub: "issip",
			},
			want: 4,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.str, tt.args.sub); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
