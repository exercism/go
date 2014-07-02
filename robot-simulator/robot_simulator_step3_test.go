// +build step3 !step1,!step2

package robot

import "testing"

// Step 3 has three major changes:
//
// *  Robots run scripts rather than respond to individual commands.
// *  A log channel allows robots and the room to log messages.
// *  The room allows multiple robots to exist and operate concurrently.
//
// Step 3 API:
//
//    Robot3(name, script string, action chan Action3, log chan string)
//    Room3(extent Rect, robots []Place,
//       action chan Action3, report chan []Place, log chan string)
//
// Again, you define Action3.
//
// For the final position report sent from Room3, you can return the same slice
// recieved from the robots channel, just with updated positions and directions.
//
// Messages must be sent on the log channel for
// *  A robot without a name
// *  Duplicate robot names
// *  Robots placed at the same place
// *  A robot placed outside of the room
// *  An undefined command in a script
// *  An action from an unknown robot
// *  A robot attempting to advance into a wall
// *  A robot attempting to advance into another robot

var testOneRobot = []struct {
	cmd byte
	DirAt
	nMsg int
}{
	// (test2 data with message counts added)
	{' ', DirAt{N, Pos{1, 1}}, 0},
	{'A', DirAt{N, Pos{1, 2}}, 0},
	{'R', DirAt{E, Pos{1, 2}}, 0},
	{'A', DirAt{E, Pos{2, 2}}, 0},
	{'L', DirAt{N, Pos{2, 2}}, 0},
	{'L', DirAt{W, Pos{2, 2}}, 0},
	{'L', DirAt{S, Pos{2, 2}}, 0},
	{'A', DirAt{S, Pos{2, 1}}, 0},
	{'R', DirAt{W, Pos{2, 1}}, 0},
	{'A', DirAt{W, Pos{1, 1}}, 0},
	{'A', DirAt{W, Pos{1, 1}}, 1}, // bump W wall
	{'L', DirAt{S, Pos{1, 1}}, 1},
	{'A', DirAt{S, Pos{1, 1}}, 2}, // bump S wall
	{'L', DirAt{E, Pos{1, 1}}, 2},
	{'A', DirAt{E, Pos{2, 1}}, 2},
	{'A', DirAt{E, Pos{2, 1}}, 3}, // bump E wall
	{'L', DirAt{N, Pos{2, 1}}, 3},
	{'A', DirAt{N, Pos{2, 2}}, 3},
	{'A', DirAt{N, Pos{2, 2}}, 4}, // bump N wall
}

func logMon(log chan string, nMsg chan int, t *testing.T) {
	n := 0
	for msg := range log {
		t.Log("Sim:", msg)
		n++
	}
	nMsg <- n
}

// Same tests as step 2; single robot, but expect log messages on wall bumps.
func TestOneStep3(t *testing.T) {
	for i := 1; i <= len(testOneRobot); i++ {
		log := make(chan string)
		nMsg := make(chan int)
		go logMon(log, nMsg, t)
		act := make(chan Action3)
		rep := make(chan []Place)
		go Room3(
			Rect{Pos{1, 1}, Pos{2, 2}},
			[]Place{{"Robbie", testOneRobot[0].DirAt}},
			act, rep, log)
		scr := ""
		for j := 1; j < i; j++ {
			scr += string(testOneRobot[j].cmd)
		}
		t.Logf("Script %q", scr)
		go Robot3("Robbie", scr, act, log)
		pls := <-rep
		lastTest := testOneRobot[i-1]
		if len(pls) != 1 {
			t.Fatalf("Got report on %d robots, want 1.", len(pls))
		}
		pl := pls[0]
		if pl.Name != "Robbie" {
			t.Fatalf(`Got report for robot %q, want report for "Robbie".`,
				pl.Name)
		}
		da := pl.DirAt
		want := lastTest.DirAt
		if da.Pos != want.Pos {
			t.Fatalf("Script %q, Pos = %v, want %v", scr, da.Pos, want.Pos)
		}
		if da.Dir != want.Dir {
			t.Fatalf("Script %q, Dir = %v, want %v", scr, da.Dir, want.Dir)
		}
		close(log)
		if n := <-nMsg; n != lastTest.nMsg {
			t.Errorf("%d sim messages logged, want %d.", n, lastTest.nMsg)
		}
	}
}

func TestNoName(t *testing.T) {
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{1, 1}, Pos{1, 1}},
		[]Place{{"", DirAt{N, Pos{1, 1}}}},
		act, rep, log)
	go Robot3("", "", act, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 1 {
		t.Fatalf("Got %d messages, want 1.", n)
	}
}

func TestSameName(t *testing.T) {
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{1, 1}, Pos{2, 1}},
		[]Place{
			{"Daryl", DirAt{N, Pos{1, 1}}},
			{"Daryl", DirAt{N, Pos{2, 1}}},
		},
		act, rep, log)
	go Robot3("Daryl", "", act, log)
	go Robot3("Daryl", "", act, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 1 {
		t.Fatalf("Got %d messages, want 1.", n)
	}
}

func TestSamePosition(t *testing.T) {
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{1, 1}, Pos{100, 100}},
		[]Place{
			{"Matter", DirAt{N, Pos{23, 47}}},
			{"Antimatter", DirAt{N, Pos{23, 47}}},
		},
		act, rep, log)
	go Robot3("Matter", "", act, log)
	go Robot3("Animatter", "", act, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 1 {
		t.Fatalf("Got %d messages, want 1.", n)
	}
}

func TestOutsideRoom(t *testing.T) {
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{1, 1}, Pos{1, 1}},
		[]Place{{"Elvis", DirAt{N, Pos{2, 3}}}},
		act, rep, log)
	go Robot3("Elvis", "", act, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 1 {
		t.Fatalf("Got %d messages, want 1.", n)
	}
}

func TestBadCommand(t *testing.T) {
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{0, 0}, Pos{0, 99}},
		[]Place{{"Vgr", DirAt{N, Pos{0, 99}}}},
		act, rep, log)
	go Robot3("Vgr", "RET", act, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 1 {
		t.Fatalf("Got %d messages, want 1.", n)
	}
}

func TestBadRobot(t *testing.T) {
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{0, 0}, Pos{0, 0}},
		[]Place{{"Data", DirAt{N, Pos{0, 0}}}},
		act, rep, log)
	go Robot3("Lore", "A", act, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 1 {
		t.Fatalf("Got %d messages, want 1.", n)
	}
}

func TestThree(t *testing.T) { // no bumping
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	go Robot3("clutz", "LAAARALA", act, log)
	go Robot3("sphero", "RRAAAAALA", act, log)
	go Robot3("roomba", "LAAARRRALLLL", act, log)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{-10, -10}, Pos{15, 10}},
		[]Place{
			{"clutz", DirAt{N, Pos{0, 0}}},
			{"sphero", DirAt{E, Pos{2, -7}}},
			{"roomba", DirAt{S, Pos{8, 4}}},
		},
		act, rep, log)
	pls := <-rep
	if len(pls) != 3 {
		t.Fatalf("Got report on %d robots, want 3.", len(pls))
	}
exp:
	for _, exp := range []Place{
		{"clutz", DirAt{W, Pos{-4, 1}}},
		{"sphero", DirAt{S, Pos{-3, -8}}},
		{"roomba", DirAt{N, Pos{11, 5}}},
	} {
		for _, pl := range pls {
			if pl.Name != exp.Name {
				continue
			}
			if pl.DirAt.Pos != exp.DirAt.Pos {
				t.Fatalf("%s at %v, want %v",
					pl.Name, pl.DirAt.Pos, exp.DirAt.Pos)
			}
			if pl.DirAt.Dir != exp.DirAt.Dir {
				t.Fatalf("%s facing %v, want %v",
					pl.Name, pl.DirAt.Dir, exp.DirAt.Dir)
			}
			continue exp
		}
		t.Fatalf("Missing %s", exp.Name)
	}
}

func TestBattle(t *testing.T) {
	log := make(chan string)
	nMsg := make(chan int)
	go logMon(log, nMsg, t)
	act := make(chan Action3)
	go Robot3("Toro", "AAAAAAA", act, log)
	go Robot3("Phere", "AAAAAAA", act, log)
	rep := make(chan []Place)
	go Room3(
		Rect{Pos{1, 1}, Pos{9, 9}},
		[]Place{
			{"Toro", DirAt{E, Pos{1, 5}}},
			{"Phere", DirAt{W, Pos{9, 5}}},
		},
		act, rep, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 7 {
		t.Fatalf("Got %d messages, want 7.", n)
	}
}
