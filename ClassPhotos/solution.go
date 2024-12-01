package ClassPhotos

import "slices"

func ClassPhotos_V1(redStudents, blueStudents []int) bool {
	var taller, smaller []int

	slices.Sort(redStudents)
	slices.Sort(blueStudents)

	if redStudents[len(redStudents)-1] > blueStudents[len(blueStudents)-1] {
		taller = redStudents
		smaller = blueStudents
	} else {
		taller = blueStudents
		smaller = redStudents
	}

	for idx, _ := range taller {
		if taller[idx] < smaller[idx] {
			return false
		}
	}

	return true
}
