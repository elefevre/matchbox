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
	return reflect.ValueOf(coinsBelow4CoinsInAColumnType(r.Int31n(int32(numberOfRows - 3))))
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
	return reflect.ValueOf(startColumnFor4CoinsInARowType(r.Int31n(int32(numberOfColumns - 3))))
}

type coinsBelow4CoinsInARowType byte

func (coinsBelow4CoinsInARowType) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(coinsBelow4CoinsInARowType(r.Int31n(int32(numberOfRows))))
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

func TestBoard_player_wins_when_there_are_4_consecutive_coins_in_a_diagonal_starting_from_the_lower_left(t *testing.T) {
	comm := func(column startColumnFor4CoinsInARowType, numberOfOpponentCoinsBelow coinsBelow4CoinsInAColumnType, player PlayerColor) bool {
		opponent := yellow
		if player == yellow {
			opponent = red
		}

		board := &Board{}

		for i := byte(0); i < byte(numberOfOpponentCoinsBelow); i++ {
			board.InsertCoin(opponent, byte(column+0))
			board.InsertCoin(opponent, byte(column+1))
			board.InsertCoin(opponent, byte(column+2))
			board.InsertCoin(opponent, byte(column+3))
		}

		board.InsertCoin(player, byte(column+0))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(opponent, byte(column+1))
		board.InsertCoin(player, byte(column+1))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(opponent, byte(column+2))
		board.InsertCoin(opponent, byte(column+2))
		board.InsertCoin(player, byte(column+2))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(opponent, byte(column+3))
		board.InsertCoin(opponent, byte(column+3))
		board.InsertCoin(opponent, byte(column+3))
		board.InsertCoin(player, byte(column+3))

		return board.HasWon(player)
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
}

func TestBoard_player_wins_when_there_are_4_consecutive_coins_in_a_diagonal_starting_from_the_upper_right(t *testing.T) {
	comm := func(column startColumnFor4CoinsInARowType, numberOfOpponentCoinsBelow coinsBelow4CoinsInAColumnType, player PlayerColor) bool {
		opponent := yellow
		if player == yellow {
			opponent = red
		}

		board := &Board{}

		for i := byte(0); i < byte(numberOfOpponentCoinsBelow); i++ {
			board.InsertCoin(opponent, byte(column+0))
			board.InsertCoin(opponent, byte(column+1))
			board.InsertCoin(opponent, byte(column+2))
			board.InsertCoin(opponent, byte(column+3))
		}

		board.InsertCoin(opponent, byte(column+0))
		board.InsertCoin(opponent, byte(column+0))
		board.InsertCoin(opponent, byte(column+0))
		board.InsertCoin(player, byte(column+0))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(opponent, byte(column+1))
		board.InsertCoin(opponent, byte(column+1))
		board.InsertCoin(player, byte(column+1))
		if board.HasWon(player) {
			return false
		}
		board.InsertCoin(opponent, byte(column+2))
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

type randomFullBoard Board

func (randomFullBoard) Generate(r *rand.Rand, size int) reflect.Value {
	b := &Board{}
	for row := byte(0); row < numberOfRows; row++ {
		for column := byte(0); column < numberOfColumns; column++ {
			n := r.Int31n(2)
			if n == 0 {
				b.InsertCoin(yellow, column)
			} else {
				b.InsertCoin(red, column)
			}
		}

	}
	return reflect.ValueOf(randomFullBoard(*b))
}

func TestBoard_board_is_full_when_all_slots_are_populated(t *testing.T) {
	comm := func(board randomFullBoard) bool {
		b := Board(board)
		return (&b).isFull()
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
}

type randomSparseBoard Board

func (randomSparseBoard) Generate(r *rand.Rand, size int) reflect.Value {
	b := &Board{}
	for row := byte(0); row < numberOfRows; row++ {
		for column := byte(0); column < numberOfColumns; column++ {
			n := r.Int31n(3)
			if n == 0 {
				b.InsertCoin(yellow, column)
			} else if n == 1 {
				b.InsertCoin(red, column)
			}
		}

	}
	return reflect.ValueOf(randomSparseBoard(*b))
}

func TestBoard_board_is_not_full_when_some_slots_are_still_empty(t *testing.T) {
	comm := func(board randomSparseBoard) bool {
		b := Board(board)
		b2 := &b
		return !b2.isFull()
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
}

func assertNumberOfCoins(t *testing.T, expectedNumber byte, board *Board, column byte) {
	t.Helper()
	coins := board.CountCoin(column)
	if coins != expectedNumber {
		t.Errorf("expected 0 coin in column %d, but found %d", column, coins)
	}
}
