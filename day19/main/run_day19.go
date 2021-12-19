package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day19"
	"fmt"
)

func main() {
	fmt.Println("** Day 19 **")
	fileContent := common.ReadFile("day19/input.txt")
	scanners := day19.ParseData(fileContent)
	result1, result2 := day19.Run(scanners)
	fmt.Println("Part1 result : ", result1)
	fmt.Println("Part2 result : ", result2)
}
