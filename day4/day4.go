package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("** Days 4 **")
	// filePath := "day4/day4_input.test.txt"
	filePath := "day4/day4_input.txt"
	fileContent := readFile(filePath)

	draw, boards := parseDrawAndBoards(fileContent)
	fmt.Println("Part1 result : ", step1_getWinnerBoard(draw, boards))

	draw2, boards2 := parseDrawAndBoards(fileContent)
	fmt.Println("Part2 result : ", step2_getallWinnerBoards(draw2, boards2))
}

// Reading files requires checking most calls for errors. This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []string {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}

func parseDrawAndBoards(input []string) ([]int, []board) {
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

func step1_getWinnerBoard(draw []int, boards []board) string {
	for stepCount, drawnNumber := range draw {
		for index := range boards {
			boards[index] = updateBoard(boards[index], drawnNumber)
			if isBoardWinning(boards[index]) {
				fmt.Println("Winning: ", index, boards[index], stepCount)
				return fmt.Sprintf("Winning score is: %d", getBoardScore(boards[index]))
			}
		}

	}
	return "No winner :("
}

func step2_getallWinnerBoards(draw []int, boards []board) string {
	lastScore := 0
	for stepCount, drawnNumber := range draw {
		for index := 0; index < len(boards); index++ {
			boards[index] = updateBoard(boards[index], drawnNumber)
			if isBoardWinning(boards[index]) {
				lastScore = getBoardScore(boards[index])
				fmt.Println("Winning: ", boards[index], stepCount, lastScore)
				boards[index] = getNewBoard()
			}
		}
	}
	return fmt.Sprint("Last Winner score: ", lastScore)
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func removeBoardFromList(boards []board, index int) []board {
	// boards[index] = boards[len(boards)-1]
	// return boards[:len(boards)-1]
	return append(boards[:index], boards[index+1:]...)
}
