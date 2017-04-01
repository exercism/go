// Package bowling implements scoring for the game of bowling.
package bowling

import "errors"

const testVersion = 1

var (
	ErrNegativeRollIsInvalid        = errors.New("Negative roll is invalid")
	ErrPinCountExceedsPinsOnTheLane = errors.New("Pin count exceeds pins on the lane")
	ErrPrematureScore               = errors.New("Score cannot be taken until the end of the game")
	ErrCannotRollAfterGameOver      = errors.New("Cannot roll after game is over")
)

// Game records the data to track a game's progress.
type Game struct {
	rolls       [21]int // storage for the rolls (10th frame gets up to 3)
	nRolls      int     // counts the rolls accumulated.
	nFrames     int     // counts completed frames, up to 10.
	rFrameStart int     // tracks the starting roll of each frame.
}

// NewGame returns a fresh zero-valued game struct.
func NewGame() *Game {
	return &Game{}
}

// Roll records one roll for a bowling frame with 'pins' knocked down.
// An error is possible depending on pin value and previous rolls.
func (g *Game) Roll(pins int) error {
	// Validate pin count on roll.
	if pins > 10 {
		return ErrPinCountExceedsPinsOnTheLane
	}
	if pins < 0 {
		return ErrNegativeRollIsInvalid
	}
	if g.completedFrames() == 10 {
		// Already completed 10 frames.
		return ErrCannotRollAfterGameOver
	}
	// Record the roll.
	g.rolls[g.nRolls] = pins
	g.nRolls++
	if pins == 10 && g.completedFrames() < 9 {
		// Frames 1 .. 9 can be strikes with no problems.
		g.completeTheFrame()
		return nil
	}
	if g.rollsThisFrame() == 2 {
		// Have counted 2 rolls on this frame.
		if g.rawFrameScore(g.rFrameStart) > 10 {
			// Unless we have completed 9 frames, cannot count > 10 for 1st 2 rolls.
			if g.completedFrames() != 9 {
				return ErrPinCountExceedsPinsOnTheLane
			}
		}
		if g.completedFrames() < 9 {
			// Completed frames 1 .. 9 with 2 rolls.
			g.completeTheFrame()
			return nil
		}
		// For frame 10, is it complete with 2nd roll ?
		if g.rawFrameScore(g.rFrameStart) < 10 {
			// Yes, complete with 2.
			g.completeTheFrame()
		}
	} else if g.rollsThisFrame() == 3 {
		// 3 rolls on the 10th frame.
		if g.isStrike(g.rFrameStart) {
			// First was a strike.
			if !g.isStrike(g.rFrameStart + 1) {
				// Second was NOT a strike, so last 2 rolls cannot exceed 10.
				if g.strikeBonus(g.rFrameStart) > 10 {
					return ErrPinCountExceedsPinsOnTheLane
				}
			}
			if b := g.strikeBonus(g.rFrameStart); b > 10 && b < 20 {
				// Unless one of the bonuses was a strike, bonus total cannot be 11 .. 19.
				if !g.isStrike(g.rFrameStart+1) && !g.isStrike(g.rFrameStart+2) {
					return ErrPinCountExceedsPinsOnTheLane
				}
			}
		} else if !g.isSpare(g.rFrameStart) {
			// Attempt to roll 3rd ball in 10th without strike or spare.
			return ErrCannotRollAfterGameOver
		}
		// Completed frame 10 with 3 rolls.
		g.completeTheFrame()
	}

	return nil
}

// Score returns the score of the game with a potential error.
func (g *Game) Score() (int, error) {
	if g.completedFrames() != 10 {
		return 0, ErrPrematureScore
	}

	score := 0
	frameStart := 0

	for frame := 0; frame < 10; frame++ {
		if g.isStrike(frameStart) {
			score += 10 + g.strikeBonus(frameStart)
			frameStart++
		} else if g.isSpare(frameStart) {
			score += 10 + g.spareBonus(frameStart)
			frameStart += 2
		} else {
			score += g.rawFrameScore(frameStart)
			frameStart += 2
		}
	}
	return score, nil
}

func (g *Game) rollsThisFrame() int     { return g.nRolls - g.rFrameStart }
func (g *Game) completeTheFrame()       { g.nFrames++; g.rFrameStart = g.nRolls }
func (g *Game) completedFrames() int    { return g.nFrames }
func (g *Game) isStrike(f int) bool     { return g.rolls[f] == 10 }
func (g *Game) rawFrameScore(f int) int { return g.rolls[f] + g.rolls[f+1] }
func (g *Game) spareBonus(f int) int    { return g.rolls[f+2] }
func (g *Game) strikeBonus(f int) int   { return g.rolls[f+1] + g.rolls[f+2] }
func (g *Game) isSpare(f int) bool      { return (g.rolls[f] + g.rolls[f+1]) == 10 }
