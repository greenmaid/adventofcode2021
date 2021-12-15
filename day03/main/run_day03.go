package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day03"
	"fmt"
)

func main() {
	fmt.Println("** Day 3 **")
	filePath := "day03/input.txt"
	fileContent := common.ReadFile(filePath)
	fmt.Println("Part1 result : ", day3.Step1_calculateGammaAndEpsilon(fileContent))
	fmt.Println("Part2 result : ", day3.Step2_calculateLifeSupportRating(fileContent))
}
