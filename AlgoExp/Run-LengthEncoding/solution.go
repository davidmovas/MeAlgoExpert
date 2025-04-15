package Run_LengthEncoding

import (
	"fmt"
	"strings"
)

func RunLengthEncoding(source string) string {
	var builder strings.Builder
	var data = map[int32]int{}

	for _, s := range source {
		data[s]++
	}

	for key, val := range data {
		if val > 9 {
			loss := val%9 + 1
			amount := val / 9

			for amount > 0 {
				builder.WriteString(fmt.Sprintf("%d%s", 9, string(key)))
				amount--
			}

			builder.WriteString(fmt.Sprintf("%d%s", loss, string(key)))
		} else {
			builder.WriteString(fmt.Sprintf("%d%s", val, string(key)))
		}
	}

	return builder.String()
}
