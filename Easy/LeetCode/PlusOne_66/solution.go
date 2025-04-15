package PlusOne_66

import (
	"slices"
)

func v1(digits []int) []int {
	var integer int

	for _, num := range digits {
		integer = 10*integer + num
	}

	integer++

	var result []int
	var flag = true
	var number int

	for flag {
		if integer < 10 {
			flag = false
		}

		number = integer % 10
		integer = integer / 10
		result = append(result, number)
	}

	slices.Reverse(result)

	return result
}

// TODO: NOT MY SOLUTION
func v2(digits []int) []int {
	var modifier = 1

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+modifier > 9 {
			digits[i] = 0
			modifier = 1
		} else {
			digits[i] += modifier
			modifier = 0
		}
	}

	if modifier == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func v3(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	return append([]int{1}, digits...)
}
