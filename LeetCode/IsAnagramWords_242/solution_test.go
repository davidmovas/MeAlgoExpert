package IsAnagramWords_242

import (
	"reflect"
	"testing"
)

func Test_IsAnagramWords(t *testing.T) {
	testFunctions := map[string]func(s, t string) bool{
		"IsAnagramWords_V1": v1,
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
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			args: args{
				s: "rat",
				t: "car",
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
