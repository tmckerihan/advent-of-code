package main

import (
	"bufio"
	"fmt"
	"os"
)

var CHARS_TO_CHECK = []rune{'M', 'A', 'S'}

func printArray(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func horizontalChecks(row []rune, position int, direction string) bool {
	if direction == "left" {
		if position < 3 {
			return false
		}
		for i := 0; i < 3; i++ {
			if row[position-(i+1)] != CHARS_TO_CHECK[i] {
				return false
			}
		}
	} else {
		if position > len(row)-4 {
			return false
		}
		for i := 0; i < 3; i++ {
			if row[position+(i+1)] != CHARS_TO_CHECK[i] {
				return false
			}
		}
	}
	return true
}

func traversalCheck1(grid [][]rune, position []int) int {
	count := 0
	// horizontal checks
	row := grid[position[0]]
	if horizontalChecks(row, position[1], "left") {
		count++
	}
	if horizontalChecks(row, position[1], "right") {
		count++
	}

	// vertical checks
	var column []rune
	for i := 0; i < len(grid); i++ {
		column = append(column, grid[i][position[1]])
	}
	if horizontalChecks(column, position[0], "left") {
		count++
	}
	if horizontalChecks(column, position[0], "right") {
		count++
	}

	//Diagonal checks
	// Check if we have enough space to the right and down
	if position[0] <= len(grid)-4 && position[1] <= len(row)-4 {
		// Check diagonal down-right
		diagonal := make([]rune, 4)
		for i := 0; i < 4; i++ {
			diagonal[i] = grid[position[0]+i][position[1]+i]
		}
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}
	}

	// Check if we have enough space to the right and up
	if position[0] >= 3 && position[1] <= len(row)-4 {
		// Check diagonal up-right
		diagonal := make([]rune, 4)
		for i := 0; i < 4; i++ {
			diagonal[i] = grid[position[0]-i][position[1]+i]
		}
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}
	}

	// Check if we have enough space to the left and down
	if position[0] <= len(grid)-4 && position[1] >= 3 {
		// Check diagonal down-left
		diagonal := make([]rune, 4)
		for i := 0; i < 4; i++ {
			diagonal[i] = grid[position[0]+i][position[1]-i]
		}
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}
	}

	// Check if we have enough space to the left and up
	if position[0] >= 3 && position[1] >= 3 {
		// Check diagonal up-left
		diagonal := make([]rune, 4)
		for i := 0; i < 4; i++ {
			diagonal[i] = grid[position[0]-i][position[1]-i]
		}
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}
	}
	return count
}

func traversalCheck2(grid [][]rune, position []int) int {
	count := 0
	row := grid[position[0]]

	// Check if we have enough space in all directions (need 1 space in each direction)
	if position[0] < len(grid)-1 && position[0] > 0 && position[1] < len(row)-1 && position[1] > 0 {
		// Check diagonal down-right
		diagonal := make([]rune, 4)
		diagonal[0] = '.' // Dummy value
		diagonal[1] = grid[position[0]-1][position[1]-1]
		diagonal[2] = grid[position[0]][position[1]]
		diagonal[3] = grid[position[0]+1][position[1]+1]
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}

		// Check diagonal up-right
		diagonal = make([]rune, 4)
		diagonal[0] = '.' // Dummy value
		diagonal[1] = grid[position[0]+1][position[1]-1]
		diagonal[2] = grid[position[0]][position[1]]
		diagonal[3] = grid[position[0]-1][position[1]+1]
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}

		// Check diagonal down-left
		diagonal = make([]rune, 4)
		diagonal[0] = '.' // Dummy value
		diagonal[1] = grid[position[0]-1][position[1]+1]
		diagonal[2] = grid[position[0]][position[1]]
		diagonal[3] = grid[position[0]+1][position[1]-1]
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}

		// Check diagonal up-left
		diagonal = make([]rune, 4)
		diagonal[0] = '.' // Dummy value
		diagonal[1] = grid[position[0]+1][position[1]+1]
		diagonal[2] = grid[position[0]][position[1]]
		diagonal[3] = grid[position[0]-1][position[1]-1]
		if horizontalChecks(diagonal, 0, "right") {
			count++
		}
	}

	if count >= 2 {
		count = 1
	} else {
		count = 0
	}
	return count
}

func txtToArray(txtFile string) [][]rune {
	file, err := os.Open(txtFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return grid
}

func puzzle1(grid [][]rune) {
	count := 0
	for rowIdx, row := range grid {
		for colIdx, val := range row {
			if val == 'X' {
				coords := []int{rowIdx, colIdx}
				count += traversalCheck1(grid, coords)
			}
		}
	}
	fmt.Println("Count:", count)
}

func puzzle2(grid [][]rune) {
	count := 0
	for rowIdx, row := range grid {
		for colIdx, val := range row {
			if val == 'A' {
				coords := []int{rowIdx, colIdx}
				count += traversalCheck2(grid, coords)
			}
		}
	}
	fmt.Println("Count:", count)
}

func main() {
	grid := txtToArray("input.txt")
	// printArray(grid)
	puzzle1(grid)
	puzzle2(grid)
}
