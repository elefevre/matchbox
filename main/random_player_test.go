package main

import "testing"

func TestRandomPlayer_plays_in_one_of_the_7_columns(t *testing.T) {
	player := randomPlayer{}

	nextPlay := player.NextPlay()
	if nextPlay < 0 || nextPlay >= 7 {
		t.Errorf("expected to play one of the 7 columns, but got an %d", nextPlay)
	}
}


func TestRandomPlayer_plays_in_a_column_with_free_slots(t *testing.T) {
	player := randomPlayer{}

	nextPlay := player.NextPlay()
	if nextPlay < 0 || nextPlay >= 7 {
		t.Errorf("expected to play one of the 7 columns, but got an %d", nextPlay)
	}
}
