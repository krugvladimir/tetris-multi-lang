package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const width = 10
const height = 20

var shapes = [][][]int{
	{{1, 1, 1, 1}},
	{{1, 1}, {1, 1}},
	{{0, 1, 0}, {1, 1, 1}},
	{{1, 0, 0}, {1, 1, 1}},
	{{0, 0, 1}, {1, 1, 1}},
	{{0, 1, 1}, {1, 1, 0}},
	{{1, 1, 0}, {0, 1, 1}},
}

type Tetris struct {
	field     [][]int
	score     int
	gameOver  bool
	piece     [][]int
	pieceX    int
	pieceY    int
}

func NewTetris() *Tetris {
	t := &Tetris{
		field: make([][]int, height),
		score: 0,
		gameOver: false,
	}
	for i := range t.field {
		t.field[i] = make([]int, width)
	}
	t.newPiece()
	return t
}

func (t *Tetris) newPiece() {
	idx := rand.Intn(len(shapes))
	t.piece = make([][]int, len(shapes[idx]))
	for i := range shapes[idx] {
		t.piece[i] = make([]int, len(shapes[idx][i]))
		copy(t.piece[i], shapes[idx][i])
	}
	t.pieceX = width/2 - len(t.piece[0])/2
	t.pieceY = 0
	if t.collision() {
		t.gameOver = true
	}
}

func (t *Tetris) collision() bool {
	for y, row := range t.piece {
		for x, cell := range row {
			if cell == 0 {
				continue
			}
			fx := t.pieceX + x
			fy := t.pieceY + y
			if fx < 0 || fx >= width || fy >= height || (fy >= 0 && t.field[fy][fx] != 0) {
				return true
			}
		}
	}
	return false
}

func (t *Tetris) merge() {
	for y, row := range t.piece {
		for x, cell := range row {
			if cell != 0 {
				t.field[t.pieceY+y][t.pieceX+x] = 1
			}
		}
	}
	t.clearLines()
	t.newPiece()
}

func (t *Tetris) clearLines() {
	lines := 0
	for y := height - 1; y >= 0; {
		full := true
		for x := 0; x < width; x++ {
			if t.field[y][x] == 0 {
				full = false
				break
			}
		}
		if full {
			// сдвиг вниз
			for i := y; i > 0; i-- {
				t.field[i] = t.field[i-1]
			}
			t.field[0] = make([]int, width)
			lines++
		} else {
			y--
		}
	}
	t.score += lines * 100
}

func (t *Tetris) move(dx, dy int) bool {
	t.pieceX += dx
	t.pieceY += dy
	if t.collision() {
		t.pieceX -= dx
		t.pieceY -= dy
		if dy == 1 {
			t.merge()
		}
		return false
	}
	return true
}

func (t *Tetris) rotate() {
	rows := len(t.piece)
	cols := len(t.piece[0])
	rotated := make([][]int, cols)
	for i := range rotated {
		rotated[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = t.piece[i][j]
		}
	}
	oldPiece := t.piece
	t.piece = rotated
	if t.collision() {
		t.piece = oldPiece
	}
}

func clearScreen() {
	cmd := "clear"
	if runtime.GOOS == "windows" {
		cmd = "cls"
	}
	c := exec.Command(cmd)
	c.Stdout = os.Stdout
	c.Run()
}

func (t *Tetris) draw() {
	clearScreen()
	fmt.Println("Score:", t.score)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := t.field[y][x] != 0
			if !cell && t.pieceY <= y && y < t.pieceY+len(t.piece) &&
				t.pieceX <= x && x < t.pieceX+len(t.piece[0]) &&
				t.piece[y-t.pieceY][x-t.pieceX] != 0 {
				cell = true
			}
			if cell {
				fmt.Print("[]")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	if t.gameOver {
		fmt.Println("GAME OVER")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewTetris()
	reader := bufio.NewReader(os.Stdin)
	for !game.gameOver {
		game.draw()
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "a":
			game.move(-1, 0)
		case "d":
			game.move(1, 0)
		case "s":
			game.move(0, 1)
		case "w":
			game.rotate()
		case "q":
			game.gameOver = true
		default:
			game.move(0, 1)
		}
	}
	game.draw()
	fmt.Println("Final score:", game.score)
}
