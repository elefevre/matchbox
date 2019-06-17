package main

import (
	"testing"
	"testing/quick"
)

func TestReinforcedPlayer_plays_in_one_of_the_7_columns(t *testing.T) {
	comm := func(color PlayerColor) bool {
		player := newReinforcedPlayer(color)

		board := &Board{}

		nextPlay := player.NextPlay(*board)
		return nextPlay >= 0 && nextPlay < numberOfColumns
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
}


func TestReinforcedPlayer_does_not_play_in_a_full_column(t *testing.T) {
	comm := func(columnToFill columnType, color PlayerColor) bool {
		player := newReinforcedPlayer(color)
		opponent := yellow
		if color == yellow {
			opponent = red
		}

		board := &Board{}
		board.InsertCoin(opponent, byte(columnToFill))
		board.InsertCoin(opponent, byte(columnToFill))
		board.InsertCoin(opponent, byte(columnToFill))
		board.InsertCoin(opponent, byte(columnToFill))
		board.InsertCoin(opponent, byte(columnToFill))
		board.InsertCoin(opponent, byte(columnToFill))
		board.InsertCoin(opponent, byte(columnToFill))

		nextPlay := player.NextPlay(*board)
		return nextPlay != byte(columnToFill)
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}

}