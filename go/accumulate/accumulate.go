package accumulate

const testVersion = 1

func Accumulate(strings []string, operation func(string) string) []string {
	response := make([]string, len(strings))
	for i, s := range strings {
		response[i] = operation(s)
	}

	return response
}
