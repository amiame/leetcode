package main

import (
	"fmt"
)

var n = 4

func main() {
	fmt.Printf("totalNQueens(%d) = %d\n", n, totalNQueens(n))
}

func totalNQueens(n int) int {
	return backtrackNQueen(0, 0)
}

type coord struct {
	x int
	y int
}

var placedSquares = make(map[coord]interface{})

func backtrackNQueen(row int, count int) int {
	for col := 0; col < n; col++ {
		if isNotUnderAttack(row, col) {
			placeQueen(row, col)
			if row+1 == n {
				count += 1 // solution only counts if there is one queen on each row
				printCoords()
			} else {
				count = backtrackNQueen(row+1, count)
			}
			removeQueen(row, col)
		}

	}

	return count
}

func printCoords() {
	fmt.Println("Solution:")
	for c := range placedSquares {
		fmt.Printf("\t(x,y): %d,%d\n", c.x, c.y)
	}
}

func placeQueen(x, y int) {
	placedSquares[coord{x, y}] = true
}

func removeQueen(x, y int) {
	delete(placedSquares, coord{x, y})
}

func isNotUnderAttack(x, y int) bool {
	for c := range placedSquares {
		if x == c.x || y == c.y {
			return false
		}

		dx := c.x - x
		dy := c.y - y
		m := dx / dy
		if m == 1 || m == -1 {
			return false
		}
	}

	return true
}
