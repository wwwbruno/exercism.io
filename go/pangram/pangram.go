package pangram

import "strings"

const testVersion = 1

func IsPangram(s string) bool {
	if s == "" {
		return false
	}

	ls := strings.ToLower(s)
	for i := 97; i < 123; i++ {
		if !strings.ContainsRune(ls, rune(i)) {
			return false
		}
	}

	return true
}
