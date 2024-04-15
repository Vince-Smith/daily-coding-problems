package game

import (
	"errors"
	"fmt"
	"math"
)

type PlayerNum int

const (
	One PlayerNum = iota + 1
	Two
)

const BLANK = "_"

type Game struct {
	currentPlayer PlayerNum
	board         [][]string
}

func NewGame(cols, rows int) Game {
	board := make([][]string, int(rows))
	for i := range board {
		board[i] = make([]string, int(cols))
		for j := range board[i] {
			board[i][j] = BLANK
		}
	}
	return Game{
		currentPlayer: One,
		board:         board,
	}
}

func (g *Game) GetActivePlayer() string {
	return fmt.Sprintf("Player %d", g.currentPlayer)
}

func (g *Game) NextPlayer() {
	if g.currentPlayer == One {
		g.currentPlayer = Two
	} else {
		g.currentPlayer = One
	}
}

func (g *Game) Draw() {
	rows := len(g.board[0])
	cols := len(g.board)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			drawCell(g.board[j][i])
		}
		fmt.Println()
	}
	fmt.Println()
}

func drawCell(s string) {
	fmt.Printf("|%s|", s)
}

func (g *Game) Add(col int) error {
	idx := col - 1
	if g.board[idx][0] != BLANK {
		return errors.New("column is already full")
	}

	for i := len(g.board[idx]) - 1; i >= 0; i-- {
		if g.board[idx][i] == BLANK {
			g.board[idx][i] = g.getPlayerToken()
			break
		}
	}

	return nil
}

func (g *Game) getPlayerToken() string {
	if g.currentPlayer == Two {
		return "O"
	}
	return "X"
}

func (g *Game) IsVictory(col int) bool {
	row := 0
	col = col - 1
	for i, r := range g.board[col] {
		if r == BLANK {
			continue
		}
		row = i
		break
	}

	return (g.isVerticalWin(col) ||
		g.isHorizontalWin(row) ||
		g.isLDiagonalWin(col, row) ||
		g.isRDiagonalWin(col, row))
}

func (g *Game) IsFull() bool {
	return false
}

func (g *Game) isVerticalWin(col int) bool {
	cnt := 0
	t := g.getPlayerToken()

	for _, row := range g.board[col] {
		if row == t {
			cnt++
		} else {
			cnt = 0
		}

		if cnt >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) isHorizontalWin(row int) bool {
	cnt := 0
	t := g.getPlayerToken()

	for _, col := range g.board {
		if col[row] == t {
			cnt++
		} else {
			cnt = 0
		}

		if cnt >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) isLDiagonalWin(col, row int) bool {
	cnt := 0
	t := g.getPlayerToken()

	spaceAbove := row - 1
	spaceLeft := col
	smallest := int(math.Min(float64(spaceAbove), float64(spaceLeft)))

	i := col - smallest
	j := row - smallest

	for i >= 0 && j >= 0 && i < len(g.board) && j < len(g.board[i]) {
		if g.board[i][j] == t {
			cnt++
		} else {
			cnt = 0
		}

		if cnt >= 4 {
			return true
		}

		i++
		j++
	}

	return false
}

func (g *Game) isRDiagonalWin(col, row int) bool {
	cnt := 0
	t := g.getPlayerToken()

	spaceBelow := (len(g.board[0]) - 1) - row
	spaceLeft := col
	smallest := int(math.Min(float64(spaceBelow), float64(spaceLeft)))

	i := col - smallest
	j := row + smallest
	fmt.Printf("start [%d, %d]\n", i, j)

	for i >= 0 && j >= 0 && i < len(g.board) && j < len(g.board[i]) {
		fmt.Printf("INSPECTING [%d, %d]\n", i, j)
		if g.board[i][j] == t {
			cnt++
		} else {
			cnt = 0
		}

		if cnt >= 4 {
			return true
		}

		i++
		j--
	}

	return false
}
