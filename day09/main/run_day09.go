package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day09"
	"fmt"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day9")
	fmt.Println("** Day 9 **")
	filePath := "day09/input.txt"
	fileContent := common.ReadFile(filePath)
	grid := day9.ParseDataToMapGrid(fileContent)
	lowPoints, count1 := day9.Step1_FindLowPoints(grid)
	fmt.Println("Part1 result : ", count1)
	count2 := day9.Step2_FindBassins(grid, lowPoints)
	fmt.Println("Part2 result : ", count2)

}
