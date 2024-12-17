package CommonCharacters

import (
	"testing"
)

func Test_CommonCharacters(t *testing.T) {
	testFunctions := map[string]func(sets []string) []string{
		"CommonCharacters_V1": CommonCharacters,
	}

	type args struct {
		sets []string
	}

	tests := []struct {
		args args
		want map[string]struct{}
	}{
		{
			args: args{
				sets: []string{"abc", "bcd", "cbaccd"},
			},
			want: map[string]struct{}{
				"b": {},
				"c": {},
			},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.sets); !compare(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}

func compare(origin []string, want map[string]struct{}) bool {
	for _, c := range origin {
		if _, ok := want[c]; !ok {
			return false
		}
	}

	return true
}
