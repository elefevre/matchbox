package main

import "fmt"

type Player interface {
	NextPlay(Board) byte
}

type PlayerColor byte

var empty PlayerColor = 0
var red PlayerColor = 1
var yellow PlayerColor = 2

func (p *PlayerColor) String() string {
	if *p == red {
		return "red"
	}
	if *p == yellow {
		return "yellow"
	}
	return "<empty>"
}

func (p *PlayerColor) ShortString() string {
	if *p == red {
		return "X"
	}
	if *p == yellow {
		return "O"
	}
	return "."
}

const numberOfColumns byte = 7
const numberOfRows byte = 6

type Board struct {
	coins [numberOfColumns][numberOfRows]PlayerColor
}

func (b *Board) InsertCoin(coin PlayerColor, column byte) error {
	for row := byte(0); row < numberOfRows; row++ {
		if b.coins[column][row] == empty {
			b.coins[column][row] = coin
			return nil
		}
	}
	return fmt.Errorf("column %d is full", column)
}

func (b *Board) CountCoin(column byte) byte {
	for row := byte(0); row < numberOfRows; row++ {
		if b.coins[column][row] == empty {
			return row
		}
	}
	return numberOfRows
}

func (b *Board) HasWon(color PlayerColor) bool {
	for column := byte(0); column < numberOfColumns; column++ {
		for row := byte(0); row < numberOfRows-3; row++ {
			if b.coins[column][row+0] == color && b.coins[column][row+1] == color && b.coins[column][row+2] == color && b.coins[column][row+3] == color {
				return true
			}
		}
	}

	for row := byte(0); row < numberOfRows; row++ {
		for column := byte(0); column < numberOfColumns-3; column++ {
			if b.coins[column+0][row] == color && b.coins[column+1][row] == color && b.coins[column+2][row] == color && b.coins[column+3][row] == color {
				return true
			}
		}
	}

	for row := byte(0); row < numberOfRows-3; row++ {
		for column := byte(0); column < numberOfColumns-3; column++ {
			if b.coins[column+0][row+0] == color && b.coins[column+1][row+1] == color && b.coins[column+2][row+2] == color && b.coins[column+3][row+3] == color {
				return true
			}
		}
	}

	for row := byte(0); row < numberOfRows-3; row++ {
		for column := byte(0); column < numberOfColumns-3; column++ {
			if b.coins[column+3][row+0] == color && b.coins[column+2][row+1] == color && b.coins[column+1][row+2] == color && b.coins[column+0][row+3] == color {
				return true
			}
		}
	}

	return false
}

func (b *Board) String() string {
	return fmt.Sprintf(`
%s%s%s%s%s%s%s
%s%s%s%s%s%s%s
%s%s%s%s%s%s%s
%s%s%s%s%s%s%s
%s%s%s%s%s%s%s
%s%s%s%s%s%s%s
`,
		b.coins[0][5].ShortString(), b.coins[1][5].ShortString(), b.coins[2][5].ShortString(), b.coins[3][5].ShortString(), b.coins[4][5].ShortString(), b.coins[5][5].ShortString(), b.coins[6][5].ShortString(),
		b.coins[0][4].ShortString(), b.coins[1][4].ShortString(), b.coins[2][4].ShortString(), b.coins[3][4].ShortString(), b.coins[4][4].ShortString(), b.coins[5][4].ShortString(), b.coins[6][4].ShortString(),
		b.coins[0][3].ShortString(), b.coins[1][3].ShortString(), b.coins[2][3].ShortString(), b.coins[3][3].ShortString(), b.coins[4][3].ShortString(), b.coins[5][3].ShortString(), b.coins[6][3].ShortString(),
		b.coins[0][2].ShortString(), b.coins[1][2].ShortString(), b.coins[2][2].ShortString(), b.coins[3][2].ShortString(), b.coins[4][2].ShortString(), b.coins[5][2].ShortString(), b.coins[6][2].ShortString(),
		b.coins[0][1].ShortString(), b.coins[1][1].ShortString(), b.coins[2][1].ShortString(), b.coins[3][1].ShortString(), b.coins[4][1].ShortString(), b.coins[5][1].ShortString(), b.coins[6][1].ShortString(),
		b.coins[0][0].ShortString(), b.coins[1][0].ShortString(), b.coins[2][0].ShortString(), b.coins[3][0].ShortString(), b.coins[4][0].ShortString(), b.coins[5][0].ShortString(), b.coins[6][0].ShortString(),
	)
}

func (b *Board) isFull() bool {
	for column := byte(0); column < numberOfColumns; column++ {
		for row := byte(0); row < numberOfRows; row++ {
			if b.coins[column][row] == empty {
				return false
			}
		}
	}

	return true
}
