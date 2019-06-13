package main

import (
	"fmt"
	"os"
)

func main() {
	yellowPlayer := newRandomPlayer(yellow)

	redPlayer := newRandomPlayer(red)

	yellowWins := 0
	redWins := 0
	draws := 0
	for i := 0; i < 10000; i++ {
		winner, draw := playGame(yellowPlayer, redPlayer)
		if draw {
			draws++
		}
		if winner == yellowPlayer {
			yellowWins++
		}
		if winner == redPlayer {
			redWins++
		}
	}

	fmt.Println("yellow won", yellowWins, "times")
	fmt.Println("red won", redWins, "times")
	fmt.Println("draws", draws, "times")
}

func playGame(yellowPlayer Player, redPlayer Player) (winner Player, draw bool) {
	board := &Board{}
	for ; ; {
		var nextPlay byte

		nextPlay = yellowPlayer.NextPlay(*board)
		err := board.InsertCoin(yellow, nextPlay)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if board.HasWon(yellow) {
			return yellowPlayer, false
		}
		if board.isFull() {
			return nil, true
		}

		nextPlay = redPlayer.NextPlay(*board)
		err = board.InsertCoin(red, nextPlay)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if board.HasWon(red) {
			return redPlayer, false
		}
		if board.isFull() {
			return nil, true
		}
	}
}
