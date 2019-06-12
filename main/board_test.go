package main

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func TestBoard_starts_empty(t *testing.T) {
	board := &Board{}

	assertNumberOfCoins(t, 0, board, 0)
	assertNumberOfCoins(t, 0, board, 1)
	assertNumberOfCoins(t, 0, board, 2)
	assertNumberOfCoins(t, 0, board, 3)
	assertNumberOfCoins(t, 0, board, 4)
	assertNumberOfCoins(t, 0, board, 5)
	assertNumberOfCoins(t, 0, board, 6)
}

func TestBoard_can_add_a_coin(t *testing.T) {
	board := &Board{}
	board.InsertCoin(red, 0)

	assertNumberOfCoins(t, 1, board, 0)
}

func TestBoard_can_add_coins_in_different_columns(t *testing.T) {
	board := &Board{}
	board.InsertCoin(red, 0)
	board.InsertCoin(red, 1)

	assertNumberOfCoins(t, 1, board, 0)
	assertNumberOfCoins(t, 1, board, 1)
}

func Add(a, b int) int {
	return a / b
}

func (PlayerColor) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(PlayerColor(1 + r.Int31n(2)))
}

type columnType byte

func (columnType) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(columnType(r.Int31n(7)))
}

type coinsBelowType byte

func (coinsBelowType) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(coinsBelowType(r.Int31n(int32(numberOfRows - 4))))
}

func TestBoard_red_wins_when_there_are_4_red_consecutive_coins_in_a_column(t *testing.T) {
	comm := func(column columnType, numberOfOpponentCoinsBelow coinsBelowType, player PlayerColor) bool {
		opponent := yellow
		if player == yellow {
			opponent = red
		}

		board := &Board{}

		for i := byte(0); i < byte(numberOfOpponentCoinsBelow); i++ {
			board.InsertCoin(opponent, byte(column))
		}

		board.InsertCoin(player, byte(column))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(player, byte(column))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(player, byte(column))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(player, byte(column))

		return board.HasWon(player)
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
}

func TestBoard_red_wins_when_she_plays_4_coins_in_the_second_row_starting_from_first_column(t *testing.T) {
	board := &Board{}

	board.InsertCoin(red, 0)
	board.InsertCoin(yellow, 1)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(red, 3)

	board.InsertCoin(red, 0)
	assertNotWin(t, board, red)
	board.InsertCoin(red, 1)
	assertNotWin(t, board, red)
	board.InsertCoin(red, 2)
	assertNotWin(t, board, red)
	board.InsertCoin(red, 3)

	assertWin(t, board, red)
}

func TestBoard_red_wins_when_she_plays_4_coins_in_a_diagonal_from_left_bottom_corner(t *testing.T) {
	board := &Board{}

	board.InsertCoin(red, 0)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 1)
	board.InsertCoin(red, 1)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(red, 2)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 3)
	board.InsertCoin(yellow, 3)
	board.InsertCoin(yellow, 3)
	board.InsertCoin(red, 3)

	assertWin(t, board, red)
}

func TestBoard_red_wins_when_she_plays_4_coins_in_a_diagonal_from_left_bottom_corner_starting_from_the_middle(t *testing.T) {
	board := &Board{}

	board.InsertCoin(yellow, 3)
	board.InsertCoin(yellow, 3)
	board.InsertCoin(yellow, 3)
	board.InsertCoin(red, 3)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(red, 2)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 1)
	board.InsertCoin(red, 1)
	assertNotWin(t, board, red)
	board.InsertCoin(red, 0)

	assertWin(t, board, red)
}

func TestBoard_red_wins_when_she_plays_4_coins_in_the_second_row_starting_from_first_column_on_2nd_row(t *testing.T) {
	board := &Board{}

	board.InsertCoin(red, 0)
	board.InsertCoin(yellow, 1)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(red, 3)

	board.InsertCoin(red, 0)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 1)
	board.InsertCoin(red, 1)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(yellow, 2)
	board.InsertCoin(red, 2)
	assertNotWin(t, board, red)
	board.InsertCoin(yellow, 3)
	board.InsertCoin(yellow, 3)
	board.InsertCoin(yellow, 3)
	assertNotWin(t, board, red)
	board.InsertCoin(red, 3)

	assertWin(t, board, red)
}

func assertWin(t *testing.T, board *Board, color PlayerColor) {
	t.Helper()
	won := board.HasWon(color)
	if !won {
		t.Errorf("expected %s to win", color.String())
	}
}

func assertNotWin(t *testing.T, board *Board, color PlayerColor) {
	t.Helper()
	won := board.HasWon(color)
	if won {
		t.Errorf("expected %s not to win", color.String())
	}
}

func assertNumberOfCoins(t *testing.T, expectedNumber byte, board *Board, column byte) {
	t.Helper()
	coins := board.CountCoin(column)
	if coins != expectedNumber {
		t.Errorf("expected 0 coin in column %d, but found %d", column, coins)
	}
}
