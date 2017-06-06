package acronym

import "strings"

const testVersion = 3

func Abbreviate(s string) string {
	var acronym string
	s = strings.Replace(s, " ", "-", -1)
	pieces := strings.Split(s, "-")

	for _, w := range pieces {
		acronym += strings.ToUpper(w[:1])

	}

	return acronym
}
