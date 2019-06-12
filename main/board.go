package main

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

const numberOfColumns byte = 7
const numberOfRows byte = 6

type Board struct {
	coins [numberOfColumns][numberOfRows]PlayerColor
}

func (b *Board) InsertCoin(coin PlayerColor, column byte) *Board {
	for row := byte(0); row < numberOfRows; row++ {
		if b.coins[column][row] == empty {
			b.coins[column][row] = coin
			return b
		}
	}
	return b
}

func (b *Board) CountCoin(column byte) byte{
	for row := byte(0); row < numberOfRows; row++ {
		if b.coins[column][row] == empty {
			return row
		}
	}
	return numberOfRows
}

func (b *Board) HasWon(color PlayerColor) bool {
	for column := byte(0); column < numberOfColumns; column++ {
		countCoinsInARow := 0
		for row := byte(0); row < numberOfRows; row++ {
			if b.coins[column][row] == color {
				countCoinsInARow++
				if countCoinsInARow == 4 {
					return true
				}
			} else {
				countCoinsInARow = 0
			}
		}
	}

	for row := byte(0); row < numberOfRows; row++ {
		countCoinsInARow := 0
		for column := byte(0); column < numberOfColumns; column++ {
			if b.coins[column][row] == color {
				countCoinsInARow++
				if countCoinsInARow == 4 {
					return true
				}
			} else {
				countCoinsInARow = 0
			}
		}
	}

	for row := 0; row < 2; row++ {
		if b.coins[3][row+3] == color && b.coins[2][row+2] == color && b.coins[1][row+1] == color && b.coins[0][row+0] == color {
			return true
		}
	}

	return false
}
