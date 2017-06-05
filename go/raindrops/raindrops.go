package raindrops

import "strconv"

const testVersion = 3

func Convert(number int) string {
	var response string

	addFactorialString(&response, number, 3, "Pling")
	addFactorialString(&response, number, 5, "Plang")
	addFactorialString(&response, number, 7, "Plong")

	if response == "" {
		return strconv.Itoa(number)
	}

	return response
}

func addFactorialString(response *string, number, factor int, text string) {
	if number%factor == 0 {
		*response += text
	}
}
