package hello

const testVersion = 2

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
