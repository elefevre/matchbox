package main

import (
	"math/rand"
	"time"
)

type randomPlayer struct {
	player PlayerColor
	rand   *rand.Rand
}

func newRandomPlayer(player PlayerColor) *randomPlayer {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &randomPlayer{player: player, rand: rand}
}

func (r *randomPlayer) NextPlay(b Board) byte {
	for ; ; {
		candidateColumn := byte(r.rand.Int31n(int32(numberOfColumns)))
		err := b.InsertCoin(r.player, candidateColumn)
		if err == nil {
			return candidateColumn;
		}
	}
}

func (r *randomPlayer) Won() {
	// nothing to do here
}

func (r *randomPlayer) Lost() {
	// nothing to do here
}
