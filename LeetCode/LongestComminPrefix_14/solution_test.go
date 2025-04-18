package LongestComminPrefix_14

import (
	"reflect"
	"testing"
)

func Test_LongestCommonPrefix(t *testing.T) {
	testFunctions := map[string]func([]string) string{
		"LongestCommonPrefix_V1": v1,
	}

	type args struct {
		strings []string
	}

	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				strings: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			args: args{
				strings: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			args: args{
				strings: []string{"caw", "cad", "can"},
			},
			want: "ca",
		}, {
			args: args{
				strings: []string{"cir", "car"},
			},
			want: "c",
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.strings); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
