package CaesarCipherEncryptor

import (
	"reflect"
	"testing"
)

func Test_CaesarCipherEncryptor(t *testing.T) {
	testFunctions := map[string]func(str string, key int) string{
		"Ca_V1": CaesarCipherEncryptor,
	}

	type args struct {
		str string
		key int
	}

	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				str: "xyz",
				key: 2,
			},
			want: "zab",
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.str, tt.args.key); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
