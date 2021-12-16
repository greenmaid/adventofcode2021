package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day16"
	"fmt"
)

func main() {
	fmt.Println("** Day 16 **")
	filePath := "day16/input.txt"
	fileContent := common.ReadFile(filePath)
	packet := day16.ParseData(fileContent[0])
	result1 := packet.GetVersionSum()
	fmt.Println("Part1 result : ", result1)
	result2 := packet.Evaluate()
	fmt.Println("Part2 result : ", result2)
}
