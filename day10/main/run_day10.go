package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day10"
	"fmt"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day10")
	fmt.Println("** Day 10 **")
	filePath := "day10/input.txt"
	fileContent := common.ReadFile(filePath)
	count1 := day10.Step1_CheckData(fileContent)
	fmt.Println("Part1 result : ", count1)
	count2 := day10.Step2_CheckIncompleteData(fileContent)
	fmt.Println("Part2 result : ", count2)

}
