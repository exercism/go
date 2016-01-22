package hello

const TestVersion = 1

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
