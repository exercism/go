package bowling

const (
	pinsPerFrame      = 10
	framesPerGame     = 10
	maxRollsPerFrame  = 2
	maxRollsLastFrame = 3
	maxRolls          = (maxRollsPerFrame * (framesPerGame - 1)) + maxRollsLastFrame
)

type Game struct {
	rolls       [maxRolls]int // storage for the rolls
	nRolls      int           // counts the rolls accumulated.
	nFrames     int           // counts completed frames, up to framesPerGame.
	rFrameStart int           // tracks the starting roll of each frame.
}

func NewGame() *Game {
	panic("Please implement the NewGame function")
}

func (g *Game) Roll(pins int) error {
	panic("Please implement the Roll function")
}

func (g *Game) Score() (int, error) {
	panic("Please implement the Score function")
}
