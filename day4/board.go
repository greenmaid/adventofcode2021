package main

type board struct {
	grid     [5][5]int
	drawGrid [5][5]int
	lastDraw int
}

func getNewBoard() board {
	var new board
	return new
}

func updateBoard(board board, drawnNumber int) board {
	for x, line := range board.grid {
		for y, value := range line {
			if value == drawnNumber {
				board.drawGrid[x][y] = 1
				board.lastDraw = drawnNumber
			}
		}
	}
	return board
}

// Check if a draw (list of int) contains a specific value
func isValueDrawn(value int, draw []int) bool {
	for _, a := range draw {
		if a == value {
			return true
		}
	}
	return false
}

func isBoardWinning(board board) bool {

	// check horizontal match
	for _, line := range board.drawGrid {
		sum := 0
		for _, value := range line {
			sum += value
		}
		if sum == 5 {
			return true
		}
	}
	// check vertical match
	for y := 0; y < 5; y++ {
		sum := 0
		for x := 0; x < 5; x++ {
			sum += board.drawGrid[x][y]
		}
		if sum == 5 {
			return true
		}
	}
	return false
}

func getBoardScore(board board) int {
	if isBoardWinning(board) {
		// sum unmarked values
		sum := 0
		for x := 0; x<5; x++{
			for y:=0;y<5;y++ {
				if board.drawGrid[x][y] == 0 {
					sum += board.grid[x][y]
				}
			}
		}
		return sum * board.lastDraw
	}
	return 0
}
