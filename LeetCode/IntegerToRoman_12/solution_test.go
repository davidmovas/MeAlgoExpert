package IntegerToRoman_12

import (
	"reflect"
	"testing"
)

func Test_RomanToInteger(t *testing.T) {
	testFunctions := map[string]func(num int) string{
		"RomanToInteger_V1": v1,
	}

	type args struct {
		num int
	}

	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				num: 3749,
			},
			want: "MMMDCCXLIX",
		},
		{
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.num); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
