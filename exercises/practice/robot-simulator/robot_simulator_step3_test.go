//go:build step3 || (!step1 && !step2)

package robot

import "testing"

var testOneRobot = []struct {
	cmd byte
	Step2Robot
	nMsg int
}{
	// (test2 data with message counts added)
	{' ', Step2Robot{N, Pos{1, 1}}, 0},
	{'A', Step2Robot{N, Pos{1, 2}}, 0},
	{'R', Step2Robot{E, Pos{1, 2}}, 0},
	{'A', Step2Robot{E, Pos{2, 2}}, 0},
	{'L', Step2Robot{N, Pos{2, 2}}, 0},
	{'L', Step2Robot{W, Pos{2, 2}}, 0},
	{'L', Step2Robot{S, Pos{2, 2}}, 0},
	{'A', Step2Robot{S, Pos{2, 1}}, 0},
	{'R', Step2Robot{W, Pos{2, 1}}, 0},
	{'A', Step2Robot{W, Pos{1, 1}}, 0},
	{'A', Step2Robot{W, Pos{1, 1}}, 1}, // bump W wall
	{'L', Step2Robot{S, Pos{1, 1}}, 1},
	{'A', Step2Robot{S, Pos{1, 1}}, 2}, // bump S wall
	{'L', Step2Robot{E, Pos{1, 1}}, 2},
	{'A', Step2Robot{E, Pos{2, 1}}, 2},
	{'A', Step2Robot{E, Pos{2, 1}}, 3}, // bump E wall
	{'L', Step2Robot{N, Pos{2, 1}}, 3},
	{'A', Step2Robot{N, Pos{2, 2}}, 3},
	{'A', Step2Robot{N, Pos{2, 2}}, 4}, // bump N wall
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
		rep := make(chan []Step3Robot)
		go Room3(
			Rect{Pos{1, 1}, Pos{2, 2}},
			[]Step3Robot{{"Robbie", testOneRobot[0].Step2Robot}},
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
		da := pl.Step2Robot
		want := lastTest.Step2Robot
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{1, 1}, Pos{1, 1}},
		[]Step3Robot{{"", Step2Robot{N, Pos{1, 1}}}},
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{1, 1}, Pos{2, 1}},
		[]Step3Robot{
			{"Daryl", Step2Robot{N, Pos{1, 1}}},
			{"Daryl", Step2Robot{N, Pos{2, 1}}},
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{1, 1}, Pos{100, 100}},
		[]Step3Robot{
			{"Matter", Step2Robot{N, Pos{23, 47}}},
			{"Antimatter", Step2Robot{N, Pos{23, 47}}},
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{1, 1}, Pos{1, 1}},
		[]Step3Robot{{"Elvis", Step2Robot{N, Pos{2, 3}}}},
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{0, 0}, Pos{0, 99}},
		[]Step3Robot{{"Vgr", Step2Robot{N, Pos{0, 99}}}},
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{0, 0}, Pos{0, 0}},
		[]Step3Robot{{"Data", Step2Robot{N, Pos{0, 0}}}},
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{-10, -10}, Pos{15, 10}},
		[]Step3Robot{
			{"clutz", Step2Robot{N, Pos{0, 0}}},
			{"sphero", Step2Robot{E, Pos{2, -7}}},
			{"roomba", Step2Robot{S, Pos{8, 4}}},
		},
		act, rep, log)
	pls := <-rep
	if len(pls) != 3 {
		t.Fatalf("Got report on %d robots, want 3.", len(pls))
	}
exp:
	for _, exp := range []Step3Robot{
		{"clutz", Step2Robot{W, Pos{-4, 1}}},
		{"sphero", Step2Robot{S, Pos{-3, -8}}},
		{"roomba", Step2Robot{N, Pos{11, 5}}},
	} {
		for _, pl := range pls {
			if pl.Name != exp.Name {
				continue
			}
			if pl.Step2Robot.Pos != exp.Step2Robot.Pos {
				t.Fatalf("%s at %v, want %v",
					pl.Name, pl.Step2Robot.Pos, exp.Step2Robot.Pos)
			}
			if pl.Step2Robot.Dir != exp.Step2Robot.Dir {
				t.Fatalf("%s facing %v, want %v",
					pl.Name, pl.Step2Robot.Dir, exp.Step2Robot.Dir)
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
	rep := make(chan []Step3Robot)
	go Room3(
		Rect{Pos{1, 1}, Pos{9, 9}},
		[]Step3Robot{
			{"Toro", Step2Robot{E, Pos{1, 5}}},
			{"Phere", Step2Robot{W, Pos{9, 5}}},
		},
		act, rep, log)
	<-rep
	close(log)
	if n := <-nMsg; n != 7 {
		t.Fatalf("Got %d messages, want 7.", n)
	}
}
