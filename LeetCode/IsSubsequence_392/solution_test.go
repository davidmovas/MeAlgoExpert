package IsSubsequence_392

import (
	"reflect"
	"testing"
)

func Test_IsSubsequence(t *testing.T) {
	testFunctions := map[string]func(string, string) bool{
		//"IsSubsequence_V1": v1,
		"IsSubsequence_V2": v2,
	}

	type args struct {
		s, t string
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
		{
			args: args{
				s: "aaaaaa",
				t: "bbaaaa",
			},
			want: false,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.s, tt.args.t); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
