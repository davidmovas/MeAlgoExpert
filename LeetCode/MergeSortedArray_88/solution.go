package MergeSortedArray_88

func MergeSortedArray(nums1 []int, nums2 []int, m, n int) []int {
	l := m + n - 1
	i := m - 1
	j := n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[l] = nums1[i]
			i--
		} else {
			nums1[l] = nums2[j]
			j--
		}
		l--
	}

	return nums1
}
