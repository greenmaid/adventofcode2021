package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day17"
	"fmt"
)

func main() {
	fmt.Println("** Day 17 **")
	filePath := "day17/input.txt"
	fileContent := common.ReadFile(filePath)
	target := day17.ParseData(fileContent[0])
	result1 := day17.Step1(target)
	fmt.Println("Part1 result : ", result1)
	result2 := day17.Step2(target)
	fmt.Println("Part2 result : ", result2)
}
