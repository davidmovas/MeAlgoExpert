package RemoveElement_27

func RemoveElement(array []int, target int) int {
	index := 0
	for i := 0; i < len(array); i++ {
		if array[i] != target {
			array[index] = array[i]
			index++
		}
	}
	return index
}
