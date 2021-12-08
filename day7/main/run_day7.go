package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day7"
	"fmt"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day7")
	fmt.Println("** Day 7 **")
	filePath := "day7/input.txt"
	fileContent := common.ReadFile(filePath)
	positions := day7.ParseDataAsPositions(fileContent)
	fmt.Println("Part1 result : ", day7.Step1_CalculateMinFuelConsuptionForAlignment(positions))
	fmt.Println("Part2 result : ", day7.Step2_CalculateMinFuelConsuptionForAlignmentWithIncresingCost(positions))
}
