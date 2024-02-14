package main

import (
	"os"

	"github.com/01-edu/z01"
)

func errMessage() {
	err := []rune("Error")
	for i := range err {
		z01.PrintRune(err[i])
	}
	z01.PrintRune('\n')
}

var field [9][9]int

func draw() {
	for _, row := range field {
		for j, _ := range row {
			z01.PrintRune(rune(row[j]) + 48)
			if j < len(row)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func countFilledCells() int {
	count := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if field[i][j] != 0 {
				count++
			}
		}
	}
	return count
}

func alreadyInVertical(x, y, value int) bool {
	for i := 0; i < 9; i++ {
		if i != y && field[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(x, y, value int) bool {
	for i := 0; i < 9; i++ {
		if i != x && field[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x, y, value int) bool {
	sx, sy := x/3*3, y/3*3
	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 3; dx++ {
			if (sy+dy) != y && (sx+dx) != x && field[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func hasDuplicateValues() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if field[i][j] != 0 {
				value := field[i][j]
				if alreadyInVertical(j, i, value) || alreadyInHorizontal(j, i, value) || alreadyInSquare(j, i, value) {
					return true
				}
			}
		}
	}
	return false
}

func canPut(x, y, value int) bool {
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

func next(x, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x, y int) bool {
	if y == 9 {
		return true
	}
	if field[y][x] != 0 {
		return solve(next(x, y))
	} else {
		for v := 1; v <= 9; v++ {
			if canPut(x, y, v) {
				field[y][x] = v
				if solve(next(x, y)) {
					return true
				}
				field[y][x] = 0
			}
		}
		return false
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		errMessage()
		return
	}
	for i, row := range args {
		if len(row) != 9 {
			errMessage()
			return
		}
		for j, char := range row {
			if char == '.' {
				field[i][j] = 0
			} else {
				if char < '1' || char > '9' {
					errMessage()
					return
				}
				field[i][j] = int(char - '0')
			}
		}
	}

	filledCellCount := countFilledCells()
	if filledCellCount < 17 || hasDuplicateValues() {
		errMessage()
		return
	}

	if solve(0, 0) {
		draw()
	} else {
		errMessage()
	}
}
