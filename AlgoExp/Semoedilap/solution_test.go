package Semoedilap

import (
	"reflect"
	"testing"
)

func Test_Semoedilap(t *testing.T) {
	testFunctions := map[string]func(words []string) [][]string{
		"Semoedilap_V1": Semoedilap,
	}

	type args struct {
		words []string
	}

	tests := []struct {
		args args
		want [][]string
	}{
		{
			args: args{
				words: []string{"diaper", "abc", "test", "cba", "repaid"},
			},
			want: [][]string{{"diaper", "repaid"}, {"abc", "cba"}},
		},
		{
			args: args{
				words: []string{"abs", "bvo", "iram", "wow", "test", "sba", "lock", "mari", "wow"},
			},
			want: [][]string{{"abs", "sba"}, {"iram", "mari"}, {"wow", "wow"}},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.words); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
