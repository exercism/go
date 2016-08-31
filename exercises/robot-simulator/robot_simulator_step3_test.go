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
	Robot
	nMsg int
}{
	// (test2 data with message counts added)
	{' ', Robot{N, Pos{1, 1}}, 0},
	{'A', Robot{N, Pos{1, 2}}, 0},
	{'R', Robot{E, Pos{1, 2}}, 0},
	{'A', Robot{E, Pos{2, 2}}, 0},
	{'L', Robot{N, Pos{2, 2}}, 0},
	{'L', Robot{W, Pos{2, 2}}, 0},
	{'L', Robot{S, Pos{2, 2}}, 0},
	{'A', Robot{S, Pos{2, 1}}, 0},
	{'R', Robot{W, Pos{2, 1}}, 0},
	{'A', Robot{W, Pos{1, 1}}, 0},
	{'A', Robot{W, Pos{1, 1}}, 1}, // bump W wall
	{'L', Robot{S, Pos{1, 1}}, 1},
	{'A', Robot{S, Pos{1, 1}}, 2}, // bump S wall
	{'L', Robot{E, Pos{1, 1}}, 2},
	{'A', Robot{E, Pos{2, 1}}, 2},
	{'A', Robot{E, Pos{2, 1}}, 3}, // bump E wall
	{'L', Robot{N, Pos{2, 1}}, 3},
	{'A', Robot{N, Pos{2, 2}}, 3},
	{'A', Robot{N, Pos{2, 2}}, 4}, // bump N wall
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
		rep := make(chan []Robot3)
		go Room3(
			Rect{Pos{1, 1}, Pos{2, 2}},
			[]Robot3{{"Robbie", testOneRobot[0].Robot}},
			act, rep, log)
		scr := ""
		for j := 1; j < i; j++ {
			scr += string(testOneRobot[j].cmd)
		}
		t.Logf("Script %q", scr)
		go StartRobot3("Robbie", scr, act, log)
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
		da := pl.Robot
		want := lastTest.Robot
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
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{1, 1}, Pos{1, 1}},
		[]Robot3{{"", Robot{N, Pos{1, 1}}}},
		act, rep, log)
	go StartRobot3("", "", act, log)
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
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{1, 1}, Pos{2, 1}},
		[]Robot3{
			{"Daryl", Robot{N, Pos{1, 1}}},
			{"Daryl", Robot{N, Pos{2, 1}}},
		},
		act, rep, log)
	go StartRobot3("Daryl", "", act, log)
	go StartRobot3("Daryl", "", act, log)
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
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{1, 1}, Pos{100, 100}},
		[]Robot3{
			{"Matter", Robot{N, Pos{23, 47}}},
			{"Antimatter", Robot{N, Pos{23, 47}}},
		},
		act, rep, log)
	go StartRobot3("Matter", "", act, log)
	go StartRobot3("Antimatter", "", act, log)
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
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{1, 1}, Pos{1, 1}},
		[]Robot3{{"Elvis", Robot{N, Pos{2, 3}}}},
		act, rep, log)
	go StartRobot3("Elvis", "", act, log)
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
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{0, 0}, Pos{0, 99}},
		[]Robot3{{"Vgr", Robot{N, Pos{0, 99}}}},
		act, rep, log)
	go StartRobot3("Vgr", "RET", act, log)
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
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{0, 0}, Pos{0, 0}},
		[]Robot3{{"Data", Robot{N, Pos{0, 0}}}},
		act, rep, log)
	go StartRobot3("Lore", "A", act, log)
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
	go StartRobot3("clutz", "LAAARALA", act, log)
	go StartRobot3("sphero", "RRAAAAALA", act, log)
	go StartRobot3("roomba", "LAAARRRALLLL", act, log)
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{-10, -10}, Pos{15, 10}},
		[]Robot3{
			{"clutz", Robot{N, Pos{0, 0}}},
			{"sphero", Robot{E, Pos{2, -7}}},
			{"roomba", Robot{S, Pos{8, 4}}},
		},
		act, rep, log)
	pls := <-rep
	if len(pls) != 3 {
		t.Fatalf("Got report on %d robots, want 3.", len(pls))
	}
exp:
	for _, exp := range []Robot3{
		{"clutz", Robot{W, Pos{-4, 1}}},
		{"sphero", Robot{S, Pos{-3, -8}}},
		{"roomba", Robot{N, Pos{11, 5}}},
	} {
		for _, pl := range pls {
			if pl.Name != exp.Name {
				continue
			}
			if pl.Robot.Pos != exp.Robot.Pos {
				t.Fatalf("%s at %v, want %v",
					pl.Name, pl.Robot.Pos, exp.Robot.Pos)
			}
			if pl.Robot.Dir != exp.Robot.Dir {
				t.Fatalf("%s facing %v, want %v",
					pl.Name, pl.Robot.Dir, exp.Robot.Dir)
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
	go StartRobot3("Toro", "AAAAAAA", act, log)
	go StartRobot3("Phere", "AAAAAAA", act, log)
	rep := make(chan []Robot3)
	go Room3(
		Rect{Pos{1, 1}, Pos{9, 9}},
		[]Robot3{
			{"Toro", Robot{E, Pos{1, 5}}},
			{"Phere", Robot{W, Pos{9, 5}}},
		},
		act, rep, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 7 {
		t.Fatalf("Got %d messages, want 7.", n)
	}
}
