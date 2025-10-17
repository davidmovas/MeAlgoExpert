package HappyNumber_202

import (
	"reflect"
	"testing"
)

func Test_HappyNumber(t *testing.T) {
	testFunctions := map[string]func(n int) bool{
		"HappyNumber_V1": v1,
	}

	type args struct {
		n int
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			args: args{
				n: 2,
			},
			want: false,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
