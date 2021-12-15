package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day08"
	"fmt"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day8")
	fmt.Println("** Day 8 **")
	filePath := "day08/input.txt"
	fileContent := common.ReadFile(filePath)
	messages := day8.ParseData(fileContent)
	count1 := day8.Step1_CountOccurrenceOfSpecialDigitsInOutput(messages)
	fmt.Println("Part1 result : ", count1)
	count2 := day8.Step2_GessDigits(messages)
	fmt.Println("Part2 result : ", count2)
}
