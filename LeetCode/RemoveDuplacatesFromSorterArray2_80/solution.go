package RemoveDuplacatesFromSorterArray_26

func RemoveDuplicatesFromSortedArray(array []int) int {
	index := 1
	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			array[index] = array[i]
			index++
		}
	}
	return index
}
