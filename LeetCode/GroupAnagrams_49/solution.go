package GroupAnagrams_49

import (
	"fmt"
	"sort"
	"strings"
)

func v1(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	groups := make(map[string][]string)
	buf := make([]int, 26)

	for _, s := range strs {
		for i := 0; i < 26; i++ {
			buf[i] = 0
		}

		for i := 0; i < len(s); i++ {
			buf[s[i]-'a']++
		}

		var b strings.Builder
		for i := 0; i < 26; i++ {
			b.WriteString(fmt.Sprint(buf[i], "#"))
		}

		key := b.String()
		groups[key] = append(groups[key], s)
	}

	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		sort.Strings(g)
		res = append(res, g)
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i][0] != res[j][0] {
			return res[i][0] < res[j][0]
		}
		if len(res[i]) != len(res[j]) {
			return len(res[i]) < len(res[j])
		}
		return strings.Join(res[i], ",") < strings.Join(res[j], ",")
	})

	return res
}
