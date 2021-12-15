package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day01"
	"fmt"
)

func main() {
	fmt.Println("** Day 1 **")
	filePath := "day01/input.txt"
	fileContent := common.ReadFileToInt(filePath)
	fmt.Println("Part1 result : ", day1.Step1_countIncreaseValues(fileContent))
	fmt.Println("Part2 result : ", day1.Step2_countIncreaseValuesByGroup(fileContent))
}
