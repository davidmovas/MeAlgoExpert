package RandomNote_383

import (
	"reflect"
	"testing"
)

func Test_RandomNote(t *testing.T) {
	testFunctions := map[string]func(note, mag string) bool{
		"RandomNote_V1": v1,
		"RandomNote_V2": v2,
	}

	type args struct {
		note, mag string
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				note: "a",
				mag:  "b",
			},
			want: false,
		},
		{
			args: args{
				note: "aa",
				mag:  "ab",
			},
			want: false,
		},
		{
			args: args{
				note: "aa",
				mag:  "aab",
			},
			want: true,
		},
		{
			args: args{
				note: "fihjjjjei",
				mag:  "hjibagacbhadfaefdjaeaebgi",
			},
			want: false,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.note, tt.args.mag); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
