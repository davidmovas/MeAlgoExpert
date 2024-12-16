package PalindromeCheck

func PalindromeCheck(s string) bool {
	head, tail := 0, len(s)-1

	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}
