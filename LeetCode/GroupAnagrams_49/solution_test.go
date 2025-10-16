package GroupAnagrams_49

import (
	"reflect"
	"testing"
)

func Test_GroupAnagrams(t *testing.T) {
	testFunctions := map[string]func([]string) [][]string{
		"GroupAnagrams_V1": v1,
	}

	type args struct {
		strs []string
	}

	tests := []struct {
		args args
		want [][]string
	}{
		{
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"bat"},
				{"nat", "tan"},
			},
		},
		{
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
