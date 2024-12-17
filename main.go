package main

import (
	"algoexpert/Easy/TwoNumberSum"
	algox "algoexpert/additional"
	"fmt"
)

func main() {
	fmt.Printf("Result: %v\n", algox.ArraySum(twoNumberSum.TwoNumberSumV1([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)) == algox.ArraySum([]int{11, -1}))
}
