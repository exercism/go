package robot

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.

func Right() {
	panic("Please implement the Right function")
}

func Left() {
	panic("Please implement the Left function")
}

func Advance() {
	panic("Please implement the Advance function")
}

func (d Dir) String() string {
	panic("Please implement the String function")
}

// Step 2
// Define Action type here.

func StartRobot(command chan Command, action chan Action) {
	panic("Please implement the StartRobot function")
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	panic("Please implement the Room function")
}

// Step 3
// Define Action3 type here.

func StartRobot3(name, script string, action chan Action3, log chan string) {
	panic("Please implement the StartRobot3 function")
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	panic("Please implement the Room3 function")
}
