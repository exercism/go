package robot

import "fmt"

// ======= Step 1
const testVersion = 2
const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) String() string {
	return "NESW"[d : d+1]
}

func Right() { Facing = (Facing + 1) % 4 }
func Left()  { Facing = (Facing + 3) % 4 }
func Advance() {
	if Facing&1 == 1 {
		X += 1 - int(Facing&2)
	} else {
		Y += 1 - int(Facing&2)
	}
}

// ======= Step 2

type Action byte

func StartRobot(cmd chan Command, act chan Action) {
	for c := range cmd {
		act <- Action(c)
	}
	close(act)
}

func Room(extent Rect, place Robot, act chan Action, report chan Robot) {
	for a := range act {
		switch a {
		case 'R':
			place.Dir = (place.Dir + 1) % 4
		case 'L':
			place.Dir = (place.Dir + 3) % 4
		case 'A':
			np := place.Pos
			if place.Dir&1 == 1 {
				np.Easting += 1 - RU(place.Dir&2)
			} else {
				np.Northing += 1 - RU(place.Dir&2)
			}
			if in(np, extent) {
				place.Pos = np
			}
		}
	}
	report <- place
}

func in(p Pos, ext Rect) bool {
	return p.Easting >= ext.Min.Easting && p.Easting <= ext.Max.Easting &&
		p.Northing >= ext.Min.Northing && p.Northing <= ext.Max.Northing
}

// ======= Step 3

type Action3 struct {
	name   string
	action byte
}

const beep = 7 // robots beep to communicate that they are done

func StartRobot3(name, script string, act chan Action3, log chan string) {
	for i := 0; i < len(script); i++ {
		act <- Action3{name, script[i]}
	}
	act <- Action3{name, beep}
}

func Room3(extent Rect, robots []Robot3, act chan Action3, rep chan []Robot3, log chan string) {
	// The function has multiple returns.  No matter what, rep <- is how we
	// communicate to the test program that the room is terminating.
	defer func() { rep <- robots }()
	nx := map[string]int{} // name index back into robots slice
	px := map[Pos]int{}    // position index back into robots slice
	for x, r := range robots {
		if r.Name == "" {
			log <- "Unnamed robot"
			return
		}
		if _, ok := nx[r.Name]; ok {
			log <- "Duplicate name"
			return
		}
		nx[r.Name] = x

		if !in(r.Robot.Pos, extent) {
			log <- "Robot placed outside room"
			return
		}
		if _, ok := px[r.Robot.Pos]; ok {
			log <- "Position occupied"
			return
		}
		px[r.Robot.Pos] = x
	}
	done := 0
	for a := range act {
		x, ok := nx[a.name]
		if !ok {
			log <- "Action by unknown robot"
			return
		}
		da := &robots[x].Robot
		switch a.action {
		case 'R':
			da.Dir = (da.Dir + 1) % 4
		case 'L':
			da.Dir = (da.Dir + 3) % 4
		case 'A':
			np := da.Pos
			if da.Dir&1 == 1 {
				np.Easting += 1 - RU(da.Dir&2)
			} else {
				np.Northing += 1 - RU(da.Dir&2)
			}
			if !in(np, extent) {
				log <- a.name + " bumps wall!"
				continue
			}
			if y, occupied := px[np]; occupied {
				log <- fmt.Sprintf("%s bumps %s!", a.name, robots[y].Name)
				continue
			}
			delete(px, da.Pos)
			px[np] = x
			da.Pos = np
		case beep:
			if done++; done == len(robots) {
				return
			}
		default:
			log <- "Undefined command"
			return
		}
	}
}
