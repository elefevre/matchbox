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

func (PlayerColor) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(PlayerColor(1 + r.Int31n(2)))
}

type columnType byte

func (columnType) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(columnType(r.Int31n(int32(numberOfColumns))))
}

type coinsBelow4CoinsInAColumnType byte

func (coinsBelow4CoinsInAColumnType) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(coinsBelow4CoinsInAColumnType(r.Int31n(int32(numberOfRows - 4))))
}

func TestBoard_player_wins_when_there_are_4_consecutive_coins_in_a_column(t *testing.T) {
	comm := func(column columnType, numberOfOpponentCoinsBelow coinsBelow4CoinsInAColumnType, player PlayerColor) bool {
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

type startColumnFor4CoinsInARowType byte

func (startColumnFor4CoinsInARowType) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(startColumnFor4CoinsInARowType(r.Int31n(int32(numberOfColumns - 4))))
}

type coinsBelow4CoinsInARowType byte

func (coinsBelow4CoinsInARowType) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(coinsBelow4CoinsInARowType(r.Int31n(int32(numberOfRows - 1))))
}

func TestBoard_player_wins_when_there_are_4_consecutive_coins_in_a_row(t *testing.T) {
	comm := func(column startColumnFor4CoinsInARowType, numberOfOpponentCoinsBelow coinsBelow4CoinsInARowType, player PlayerColor) bool {
		opponent := yellow
		if player == yellow {
			opponent = red
		}

		board := &Board{}

		for i := byte(0); i < byte(numberOfOpponentCoinsBelow); i++ {
			board.InsertCoin(opponent, byte(column))
			board.InsertCoin(opponent, byte(column+1))
			board.InsertCoin(opponent, byte(column+2))
			board.InsertCoin(opponent, byte(column+3))
		}

		board.InsertCoin(player, byte(column))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(player, byte(column+1))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(player, byte(column+2))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(player, byte(column+3))

		return board.HasWon(player)
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
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
