package greeting

const testVersion = 3

// HelloWorld says hello to given name.
// If no name is passed, it greets the world.
func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
