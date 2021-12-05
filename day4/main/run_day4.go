package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day4"
	"fmt"
)

func main() {
	fmt.Println("** Day 4 **")
	filePath := "day4/input.txt"
	fileContent := common.ReadFile(filePath)

	draw, boards := day4.ParseDrawAndBoards(fileContent)
	fmt.Println("Part1 result : ", day4.Step1_getWinnerBoard(draw, boards))

	draw2, boards2 := day4.ParseDrawAndBoards(fileContent)
	fmt.Println("Part2 result : ", day4.Step2_getAllWinnerBoards(draw2, boards2))
}
