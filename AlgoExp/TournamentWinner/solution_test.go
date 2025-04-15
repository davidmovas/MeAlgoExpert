package TournamentWinner

import (
	"reflect"
	"testing"
)

func Test_tournamentWinner_test(t *testing.T) {
	testFunctions := map[string]func(competitions [][]string, results []int) string{
		"TournamentWinner_test_V1": TournamentWinner,
	}

	type args struct {
		competitions [][]string
		results      []int
	}

	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				competitions: [][]string{
					{"HTML", "C#"},
					{"C#", "Go"},
					{"Go", "HTML"},
				},
				results: []int{0, 0, 1},
			},
			want: "Go",
		},
		{
			args: args{
				competitions: [][]string{
					{"HTML", "C#"},
					{"C#", "Go"},
					{"Go", "HTML"},
				},
				results: []int{0, 1, 0},
			},
			want: "C#",
		},
		{
			args: args{
				competitions: [][]string{
					{"HTML", "C#"},
					{"C#", "Go"},
					{"Go", "HTML"},
				},
				results: []int{1, 0, 0},
			},
			want: "HTML",
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.competitions, tt.args.results); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
