package greeting

const testVersion = 3

func HelloWorld(name string) string {

	var helloName string = "World"

	if len(name) > 0 {
		helloName = name
	}

	return "Hello, " + helloName + "!"
}
