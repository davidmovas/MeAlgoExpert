package BinarySearch

func BinarySearch(array []int, target int) int {
	head, tail := 0, len(array)-1

	for head <= tail {
		mid := (head + tail) / 2
		potential := array[mid]
		if potential == target {
			return mid
		} else if target > potential {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}

	return -1
}
