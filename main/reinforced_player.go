package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type reinforcedPlayer struct {
	color                      PlayerColor
	rand                       *rand.Rand
	beads                      map[Board][numberOfColumns]int
	boardsInCurrentGame        []Board
	columnsPlayed              []byte
	numberOfMatchBoxesNotified int
	numberOfGamesNotified      int
}

func (r *reinforcedPlayer) NextPlay(b Board) byte {
	beads, ok := r.beads[b]
	if !ok {
		// first time we find this board configuration
		beads = [numberOfColumns]int{1, 1, 1, 1, 1, 1, 1}
	}

	for ; ; {
		var candidateColumn byte
		foundCandidate := false
		for ; !foundCandidate; {
			candidateColumn = byte(r.rand.Int31n(int32(numberOfColumns)))
			if beads[candidateColumn] > 0 {
				foundCandidate = true
			}
		}
		err := b.InsertCoin(r.color, candidateColumn)
		if err != nil {
			// impossible to insert, so we discard all beads for that column
			// this should prevent it from being a candidate again
			if beads[candidateColumn] > 0 {
				beads[candidateColumn] = 0
			}
			continue
		}

		r.beads[b] = beads
		r.boardsInCurrentGame = append(r.boardsInCurrentGame, b)
		r.columnsPlayed = append(r.columnsPlayed, candidateColumn)
		return candidateColumn;
	}
}

func newReinforcedPlayer(color PlayerColor) *reinforcedPlayer {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &reinforcedPlayer{color: color, rand: rand, beads: make(map[Board][numberOfColumns]int)}
}

func (r *reinforcedPlayer) Won() {
	r.reinforce()
	r.boardsInCurrentGame = nil
	r.columnsPlayed = nil
	r.numberOfGamesNotified++

	if len(r.beads) > 1000000 + r.numberOfMatchBoxesNotified {
		r.numberOfMatchBoxesNotified = len(r.beads)

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		alloc := m.Alloc / 1024 / 1024
		totalAlloc := m.TotalAlloc / 1024 / 1024
		sys := m.Sys / 1024 / 1024
		numGC := m.NumGC / 1024 / 1024
		fmt.Printf("number of board positions saved after %d games, at %s: %d (alloc %d MB, totalAlloc %d MB, sys %d MB, numGC %d MB)\n", r.numberOfGamesNotified, time.Now(), r.numberOfMatchBoxesNotified, alloc, totalAlloc, sys, numGC)
	}
}

func (r *reinforcedPlayer) reinforce() {
	for n, boardWon := range r.boardsInCurrentGame {
		ints := r.beads[boardWon]
		ints[r.columnsPlayed[n]]++
		r.beads[boardWon] = ints
	}
}

func (r *reinforcedPlayer) Lost() {
	r.discourage()
	r.boardsInCurrentGame = nil
	r.columnsPlayed = nil
	r.numberOfGamesNotified++

	//	fmt.Println("number of board positions saved:", len(r.beads))
}

func (r *reinforcedPlayer) discourage() {
	for n, boardWon := range r.boardsInCurrentGame {
		ints := r.beads[boardWon]
		if r.beads[boardWon][r.columnsPlayed[n]] > 0 {
			ints[r.columnsPlayed[n]]--
		}
		r.beads[boardWon] = ints
	}
}
