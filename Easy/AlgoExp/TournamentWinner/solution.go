package TournamentWinner

func TournamentWinner(competitions [][]string, results []int) (winner string) {
	points := make(map[string]int)
	winner = competitions[0][0]

	for _, cc := range competitions {
		for _, c := range cc {
			points[c] = 0
		}
	}

	for indexPaar, paar := range competitions {
		home := paar[0]
		away := paar[1]

		if results[indexPaar] == 0 {
			points[away] += 3
		} else {
			points[home] += 3
		}

		if points[winner] < points[home] {
			winner = home
		} else if points[winner] < points[away] {
			winner = away
		}
	}

	return winner
}
