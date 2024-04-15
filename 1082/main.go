package main

import (
	game "Vince-Smith/connect4/internal"
	"fmt"
	"strconv"
)

const COLS = 7
const ROWS = 6

func main() {
	g := game.NewGame(COLS, ROWS)
	fmt.Println("Welcome to Connect Four!")
	fmt.Println("------------------------")
	for {
		g.Draw()
		fmt.Printf("%s:\n", g.GetActivePlayer())
		fmt.Printf("Choose a column between 1 and %d:\n", COLS)

		var s string
		_, err := fmt.Scan(&s)

		if err != nil {
			panic(err)
		}

		if !isValidInput(s) {
			fmt.Printf("%s is not a valid input. Try again!\n", s)
			continue
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Something went wrong")
			continue
		}
		err = g.Add(i)
		if err != nil {
			fmt.Println("Cannot add to this column! Try again.")
			continue
		}

		if g.IsVictory(i) {
			fmt.Println("CONGRATS!")
			fmt.Printf("Player %s is the winner!\n", g.GetActivePlayer())
			break
		}

		if g.IsFull() {
			fmt.Println("DRAW!")
		}

		g.NextPlayer()
	}
}

func isValidInput(s string) bool {
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if num < 1 || num > COLS {
		return false
	}

	return true
}
