// Package bowling implements scoring for the game of bowling.
package bowling

import "errors"

var (
	ErrNegativeRollIsInvalid        = errors.New("Negative roll is invalid")
	ErrPinCountExceedsPinsOnTheLane = errors.New("Pin count exceeds pins on the lane")
	ErrPrematureScore               = errors.New("Score cannot be taken until the end of the game")
	ErrCannotRollAfterGameOver      = errors.New("Cannot roll after game is over")
)

const (
	pinsPerFrame      = 10
	framesPerGame     = 10
	maxRollsPerFrame  = 2
	maxRollsLastFrame = 3
	maxRolls          = (maxRollsPerFrame * (framesPerGame - 1)) + maxRollsLastFrame
)

// Game records the data to track a game's progress.
type Game struct {
	rolls       [maxRolls]int // storage for the rolls
	nRolls      int           // counts the rolls accumulated.
	nFrames     int           // counts completed frames, up to framesPerGame.
	rFrameStart int           // tracks the starting roll of each frame.
}

// NewGame returns a fresh zero-valued game struct.
func NewGame() *Game {
	return &Game{}
}

// Roll records one roll for a bowling frame with 'pins' knocked down.
// An error is possible depending on pin value and previous rolls.
func (g *Game) Roll(pins int) error {
	// Validate pin count on roll.
	if pins > pinsPerFrame {
		return ErrPinCountExceedsPinsOnTheLane
	}
	if pins < 0 {
		return ErrNegativeRollIsInvalid
	}
	if g.completedFrames() == framesPerGame {
		return ErrCannotRollAfterGameOver
	}
	// Record the roll.
	g.rolls[g.nRolls] = pins
	g.nRolls++
	if pins == pinsPerFrame && g.completedFrames() < framesPerGame-1 {
		// Frames before last one can be strikes with no problems.
		g.completeTheFrame()
		return nil
	}
	if g.rollsThisFrame() == maxRollsPerFrame {
		// Have counted normal max rolls on a frame.
		if g.rawFrameScore(g.rFrameStart) > pinsPerFrame {
			// Unless we have completed all but last frame, cannot count > pinsPerFrame.
			if g.completedFrames() != framesPerGame-1 || !g.isStrike(g.rFrameStart) {
				return ErrPinCountExceedsPinsOnTheLane
			}
		}
		if g.completedFrames() < framesPerGame-1 {
			// Completed frames before last one with maxRollsPerFrame.
			g.completeTheFrame()
			return nil
		}
		// For last frame, is it complete now ?
		if g.rawFrameScore(g.rFrameStart) < pinsPerFrame {
			// Yes, complete.
			g.completeTheFrame()
		}
	} else if g.rollsThisFrame() == maxRollsLastFrame {
		// Extra roll on the last frame.
		if g.isStrike(g.rFrameStart) {
			// First was a strike.
			if !g.isStrike(g.rFrameStart + 1) {
				// Second was NOT a strike, so last 2 rolls cannot exceed pinsPerFrame.
				if g.strikeBonus(g.rFrameStart) > pinsPerFrame {
					return ErrPinCountExceedsPinsOnTheLane
				}
			}
			if b := g.strikeBonus(g.rFrameStart); b > pinsPerFrame && b < 2*pinsPerFrame {
				// Unless one of the bonuses was a strike, bonus frames too high.
				if !g.isStrike(g.rFrameStart+1) && !g.isStrike(g.rFrameStart+2) {
					return ErrPinCountExceedsPinsOnTheLane
				}
			}
		} else if !g.isSpare(g.rFrameStart) {
			// Attempt to make extra roll in last frame without strike or spare.
			return ErrCannotRollAfterGameOver
		}
		// Completed last frame.
		g.completeTheFrame()
	}

	return nil
}

// Score returns the score of the game with a potential error.
func (g *Game) Score() (int, error) {
	if g.completedFrames() != framesPerGame {
		return 0, ErrPrematureScore
	}

	score := 0
	frameStart := 0

	for frame := 0; frame < framesPerGame; frame++ {
		switch {
		case g.isStrike(frameStart):
			score += pinsPerFrame + g.strikeBonus(frameStart)
			frameStart++
		case g.isSpare(frameStart):
			score += pinsPerFrame + g.spareBonus(frameStart)
			frameStart += maxRollsPerFrame
		default:
			score += g.rawFrameScore(frameStart)
			frameStart += maxRollsPerFrame
		}
	}
	return score, nil
}

func (g *Game) rollsThisFrame() int     { return g.nRolls - g.rFrameStart }
func (g *Game) completeTheFrame()       { g.nFrames++; g.rFrameStart = g.nRolls }
func (g *Game) completedFrames() int    { return g.nFrames }
func (g *Game) isStrike(f int) bool     { return g.rolls[f] == pinsPerFrame }
func (g *Game) rawFrameScore(f int) int { return g.rolls[f] + g.rolls[f+1] }
func (g *Game) spareBonus(f int) int    { return g.rolls[f+2] }
func (g *Game) strikeBonus(f int) int   { return g.rolls[f+1] + g.rolls[f+2] }
func (g *Game) isSpare(f int) bool      { return (g.rolls[f] + g.rolls[f+1]) == pinsPerFrame }
