package main

import (
	"adventofcode2021/day2"
	"adventofcode2021/common"
	"fmt"
)

func main() {
	fmt.Println("** Day 2 **")
	filePath := "day2/input.txt"
	fileContent := common.ReadFile(filePath)
	fmt.Println("Part1 result : ", day2.Step1_followInstructions(fileContent))
	fmt.Println("Part2 result : ", day2.Step2_followInstructionsWithAim(fileContent))
}
