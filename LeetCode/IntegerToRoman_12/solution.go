package IntegerToRoman_12

import "strings"

var romanianMap = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

var romanArray = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var integerArray = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func v1(num int) string {
	var builder strings.Builder

	for i, value := range integerArray {
		if num == 0 {
			return builder.String()
		}

		if num >= value {
			times := num / value
			for range times {
				builder.WriteString(romanArray[i])
			}
			num = num % (value * times)
		}
	}

	return builder.String()
}
