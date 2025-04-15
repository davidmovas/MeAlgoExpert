package Run_LengthEncoding

import (
	"reflect"
	"testing"
)

func Test_RunLengthEncoding(t *testing.T) {
	testFunctions := map[string]func(source string) string{
		"RunLengthEncoding_V1": RunLengthEncoding,
	}

	type args struct {
		source string
	}

	tests := []struct {
		args args
		want any
	}{
		{
			args: args{
				source: "AAAAAAAAAAAABBCCCCDD",
			},
			want: "9A4A2B4C2D",
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.source); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
