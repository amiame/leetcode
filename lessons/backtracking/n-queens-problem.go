package main

import (
	"fmt"
	"log"
)

var n = 0

func exitIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var i int
	for {
		_, err := fmt.Scanf("%d", &i)
		exitIfErr(err)
		n = i
		fmt.Printf("totalNQueens(%d) = %d solutions\n", n, totalNQueens())
		printSolutions("Solutions")
		clearSolutions()
	}
}

func totalNQueens() int {
	return backtrackNQueen(0, 0)
}

type coord struct {
	x int
	y int
}

var placedSquares = make(map[coord]interface{})
var solutions = make([]map[coord]interface{}, 0)

func backtrackNQueen(row int, count int) int {
	for col := 0; col < n; col++ {
		if isNotUnderAttack(row, col) {
			placeQueen(row, col)
			if row+1 == n {
				count += 1 // solution only counts if there is one queen on each row
				//		printCoords("Solution")
				saveSolution()
			} else {
				count = backtrackNQueen(row+1, count)
			}
			removeQueen(row, col)
		}

	}

	return count
}

func clearSolutions() {
	solutions = make([]map[coord]interface{}, 0)
}

func printSolutions(title string) {
	fmt.Printf("%s:\n", title)
	for row := 0; row < n; row++ {
		for i, s := range solutions {
			if i == 0 {
				fmt.Printf("\t")
			} else {
				fmt.Printf("  ")
			}

			for col := 0; col < n; col++ {
				if _, ok := s[coord{x: row, y: col}]; ok {
					fmt.Printf("o")
				} else {
					fmt.Printf("-")
				}
			}
		}
		fmt.Println()
	}
}

func saveSolution() {
	solution := make(map[coord]interface{})
	for k, v := range placedSquares {
		solution[k] = v
	}
	solutions = append(solutions, solution)
}

func printCoords(title string) {
	fmt.Printf("%s:\n", title)
	for i := 0; i < n; i++ {
		fmt.Printf("\t")
		for j := 0; j < n; j++ {
			if _, ok := placedSquares[coord{x: i, y: j}]; ok {
				fmt.Printf("o")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
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
		mmod := dx % dy
		m := dx / dy
		if mmod == 0 && (m == 1 || m == -1) {
			return false
		}
	}

	return true
}
