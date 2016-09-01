// +build step2 !step1,!step3

package robot

import "testing"

// For step 1 you implement robot movements, but it's not much of a simulation.
// For example where in the source code is "the robot"?  Where is "the grid"?
// Where are the computations that turn robot actions into grid positions,
// in the robot or in the grid?  The physical world is different.
//
// Step 2 introduces a "room."  It seems a small addition, but we'll make
// big changes to clarify the rolls of "room", "robot", and "test program"
// and begin to clarify the physics of the simulation.  You will define Room
// and Robot as functions which the test program "brings into existence" by
// launching them as goroutines.  Information moves between test program,
// robot, and room over Go channels.
//
// Think of Room as a "physics engine," something that models and simulates
// a physical room with walls and a robot.  It should somehow model the
// coordinate space of the room, the location of the robot and the walls,
// and ensure for example that the robot doesn't walk through walls.
// We want Robot to be an agent that performs actions, but we want Room to
// maintain a coherent truth.
//
// Step 2 API:
//
// Robot(chan Command, chan Action)
// Room(extent Rect, place Robot2, act chan Action, rep chan Robot2)
//
// You get to define Action; see defs.go for other definitions.
//
// The test program creates the channels and starts both Room and Robot.
// The test program then sends commands to Robot.  When it is done sending
// commands, it closes the command channel.  Robot must accept commands and
// inform Room of actions it is attempting.  When it senses the command channel
// closing, it must shut down itself.  The room must interpret the physical
// consequences of the robot actions.  When it senses the robot shutting down,
// it sends a final report back to the test program, telling the robot's final
// position and direction.

var test2 = []struct {
	Command
	Robot2
}{
	0:  {' ', Robot2{N, Pos{1, 1}}}, // no command, this is the start Robot2
	1:  {'A', Robot2{N, Pos{1, 2}}},
	2:  {'R', Robot2{E, Pos{1, 2}}},
	3:  {'A', Robot2{E, Pos{2, 2}}},
	4:  {'L', Robot2{N, Pos{2, 2}}},
	5:  {'L', Robot2{W, Pos{2, 2}}},
	6:  {'L', Robot2{S, Pos{2, 2}}},
	7:  {'A', Robot2{S, Pos{2, 1}}},
	8:  {'R', Robot2{W, Pos{2, 1}}},
	9:  {'A', Robot2{W, Pos{1, 1}}},
	10: {'A', Robot2{W, Pos{1, 1}}}, // bump W wall
	11: {'L', Robot2{S, Pos{1, 1}}},
	12: {'A', Robot2{S, Pos{1, 1}}}, // bump S wall
	13: {'L', Robot2{E, Pos{1, 1}}},
	14: {'A', Robot2{E, Pos{2, 1}}},
	15: {'A', Robot2{E, Pos{2, 1}}}, // bump E wall
	16: {'L', Robot2{N, Pos{2, 1}}},
	17: {'A', Robot2{N, Pos{2, 2}}},
	18: {'A', Robot2{N, Pos{2, 2}}}, // bump N wall
}

func TestStep2(t *testing.T) {
	// run incrementally longer tests
	for i := 1; i <= len(test2); i++ {
		cmd := make(chan Command)
		act := make(chan Action)
		rep := make(chan Robot2)
		go StartRobot(cmd, act)
		go Room(Rect{Pos{1, 1}, Pos{2, 2}}, test2[0].Robot2, act, rep)
		for j := 1; j < i; j++ {
			cmd <- test2[j].Command
		}
		close(cmd)
		da := <-rep
		last := i - 1
		want := test2[last].Robot2
		if da.Pos != want.Pos {
			t.Fatalf("Command #%d, Pos = %v, want %v", last, da.Pos, want.Pos)
		}
		if da.Dir != want.Dir {
			t.Fatalf("Command #%d, Dir = %v, want %v", last, da.Dir, want.Dir)
		}
	}
}
