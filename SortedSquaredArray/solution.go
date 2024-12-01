package SortedSquaredArray

func SortedSquaredArrayV1(array []int) []int {
	for index, _ := range array {
		array[index] = array[index] * array[index]
	}
	return array
}
