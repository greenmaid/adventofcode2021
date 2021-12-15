package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day00"
	"fmt"
)

func main() {
	fmt.Println("** Day 00 **")
	filePath := "day00/input.txt"
	fileContent := common.ReadFile(filePath)
	data := day00.ParseData(fileContent)
	result1 := day00.Step1(data)
	fmt.Println("Part1 result : ", result1)
	result2 := day00.Step2(data)
	fmt.Println("Part2 result : ", result2)
}
