package RemoveDuplacatesFromSorterArray_26

func RemoveDuplicatesFromSortedArray(array []int) int {
	count, current := 0, 1
	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			count = 0
			array[current] = array[i]
			current++
		} else {
			count++
			if count <= 1 {
				array[current] = array[i]
				current++
			}
		}
	}
	return current
}
