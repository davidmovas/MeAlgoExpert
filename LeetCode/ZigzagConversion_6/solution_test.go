package ZigzagConversion_6

import (
	"reflect"
	"testing"
)

func Test_ZigzagConversion(t *testing.T) {
	testFunctions := map[string]func(string, int) string{
		"ZigzagConversion_V1": v1,
	}

	type args struct {
		input   string
		numRows int
	}

	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				input:   "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			args: args{
				input:   "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			args: args{
				input:   "A",
				numRows: 1,
			},
			want: "A",
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.input, tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
