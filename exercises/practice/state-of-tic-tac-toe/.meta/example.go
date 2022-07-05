package stateoftictactoe

import (
	"errors"
	"strings"
)

func StateOfTicTacToe(board []string) (string, error) {
	var xWin int
	var oWin int

	rows := board
	cols := getCols(board)
	diags := getDiags(board)
	xCount := count("X", rows)
	oCount := count("O", rows)
	if oCount > xCount {
		return "", errors.New("Wrong turn order: O started")
	}
	if xCount > oCount+1 {
		return "", errors.New("Wrong turn order: X went twice")
	}

	allCombos := append(rows, cols...)
	allCombos = append(allCombos, diags...)

	for _, row := range allCombos {
		if row == "XXX" {
			if xWin > 1 || oWin > 0 {
				return "", errors.New("Impossible board: game should have ended after the game was won")
			}
			xWin += 1
		}
		if row == "OOO" {
			if xWin > 0 || oWin > 0 {
				return "", errors.New("Impossible board: game should have ended after the game was won")
			}
			oWin = 1
		}
	}

	if xWin > 0 || oWin > 0 {
		return "win", nil
	}

	if xCount+oCount < 9 {
		return "ongoing", nil
	}

	return "draw", nil

}

func getCols(board []string) []string {
	cols := make([]string, 3)
	for _, row := range board {
		for j := 0; j < 3; j++ {
			cols[j] += string(row[j])
		}
	}
	return cols
}

func getDiags(board []string) []string {
	diags := make([]string, 2)
	for i, row := range board {
		diags[0] += string(row[i])
		diags[1] += string(row[2-i])
	}
	return diags
}

func count(c string, rows []string) int {
	num := 0
	for _, row := range rows {
		num += strings.Count(row, c)
	}
	return num
}
