package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day11"
	"fmt"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day11")
	fmt.Println("** Day 9 **")
	filePath := "day11/input.txt"
	fileContent := common.ReadFile(filePath)
	grid := day11.ParseDataToOctopusMap(fileContent)
	count1 := day11.Step1_getFlashCountAfter100Steps(grid)
	fmt.Println("Part1 result : ", count1)
	count2 := day11.Step2_FindSynchronizationStep(grid)
	fmt.Println("Part2 result : ", count2)

}
