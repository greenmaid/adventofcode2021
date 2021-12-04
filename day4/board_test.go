package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWinningBoardCheck(t *testing.T) {
	filePath := "day4_input.test.txt"
	fileContent := readFile(filePath)
	_, boards := parseDrawAndBoards(fileContent)

	for _, drawNumber := range []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24} {
		boards[2] = updateBoard(boards[2], drawNumber)
	}
	t.Log(boards[2])

	assert.True(t, isBoardWinning(boards[2]))
	assert.Equal(t, 4512, getBoardScore(boards[2]))
}
