//go:build step2 || (!step1 && !step3)

package robot

import "testing"

var test2 = []struct {
	Command
	Step2Robot
}{
	0:  {' ', Step2Robot{N, Pos{1, 1}}}, // no command, this is the start DirAt
	1:  {'A', Step2Robot{N, Pos{1, 2}}},
	2:  {'R', Step2Robot{E, Pos{1, 2}}},
	3:  {'A', Step2Robot{E, Pos{2, 2}}},
	4:  {'L', Step2Robot{N, Pos{2, 2}}},
	5:  {'L', Step2Robot{W, Pos{2, 2}}},
	6:  {'L', Step2Robot{S, Pos{2, 2}}},
	7:  {'A', Step2Robot{S, Pos{2, 1}}},
	8:  {'R', Step2Robot{W, Pos{2, 1}}},
	9:  {'A', Step2Robot{W, Pos{1, 1}}},
	10: {'A', Step2Robot{W, Pos{1, 1}}}, // bump W wall
	11: {'L', Step2Robot{S, Pos{1, 1}}},
	12: {'A', Step2Robot{S, Pos{1, 1}}}, // bump S wall
	13: {'L', Step2Robot{E, Pos{1, 1}}},
	14: {'A', Step2Robot{E, Pos{2, 1}}},
	15: {'A', Step2Robot{E, Pos{2, 1}}}, // bump E wall
	16: {'L', Step2Robot{N, Pos{2, 1}}},
	17: {'A', Step2Robot{N, Pos{2, 2}}},
	18: {'A', Step2Robot{N, Pos{2, 2}}}, // bump N wall
}

func TestStep2(t *testing.T) {
	// run incrementally longer tests
	for i := 1; i <= len(test2); i++ {
		cmd := make(chan Command)
		act := make(chan Action)
		rep := make(chan Step2Robot)
		go StartRobot(cmd, act)
		go Room(Rect{Pos{1, 1}, Pos{2, 2}}, test2[0].Step2Robot, act, rep)
		for j := 1; j < i; j++ {
			cmd <- test2[j].Command
		}
		close(cmd)
		da := <-rep
		last := i - 1
		want := test2[last].Step2Robot
		if da.Pos != want.Pos {
			t.Fatalf("Command #%d, Pos = %v, want %v", last, da.Pos, want.Pos)
		}
		if da.Dir != want.Dir {
			t.Fatalf("Command #%d, Dir = %v, want %v", last, da.Dir, want.Dir)
		}
	}
}
