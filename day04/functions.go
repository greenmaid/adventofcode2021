package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseDrawAndBoards(input []string) ([]int, []board) {
	drawStr := strings.Split(input[0], ",")
	var draw []int
	for _, value := range drawStr {
		conv, _ := strconv.Atoi(value)
		draw = append(draw, conv)
	}

	var boards []board

	currentBoard := getNewBoard()
	currentBoardIndex := 0
	for i := 2; i < len(input); i++ {

		lineNumbersStr := strings.Fields(input[i])
		// var lineNumbers []int
		for index, value := range lineNumbersStr {
			conv, _ := strconv.Atoi(value)
			currentBoard.grid[currentBoardIndex][index] = conv
		}
		currentBoardIndex += 1
		if currentBoardIndex == 5 {
			boards = append(boards, currentBoard)
			currentBoard = getNewBoard()
			currentBoardIndex = 0
			i += 1 // skip blank line
			continue
		}
	}
	return draw, boards
}

func Step1_getWinnerBoard(draw []int, boards []board) string {
	for _, drawnNumber := range draw {
		for index := range boards {
			boards[index] = updateBoard(boards[index], drawnNumber)
			if isBoardWinning(boards[index]) {
				// fmt.Println("Winning: ", index, boards[index])
				return fmt.Sprintf("Winning score is: %d", getBoardScore(boards[index]))
			}
		}

	}
	return "No winner :("
}

func Step2_getAllWinnerBoards(draw []int, boards []board) string {
	lastScore := 0
	for _, drawnNumber := range draw {
		for index := 0; index < len(boards); index++ {
			boards[index] = updateBoard(boards[index], drawnNumber)
			if isBoardWinning(boards[index]) {
				lastScore = getBoardScore(boards[index])
				// fmt.Println("Winning: ", boards[index], lastScore)
				boards[index] = getNewBoard()
			}
		}
	}
	return fmt.Sprint("Last Winner score: ", lastScore)
}
