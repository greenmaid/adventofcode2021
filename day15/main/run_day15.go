package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day15"
	"fmt"
)

func main() {
	fmt.Println("** Day 15 **")
	filePath := "day15/input.txt"
	fileContent := common.ReadFile(filePath)
	// defer common.TimeTrack(time.Now(), "day13")
	g, end := day15.ParseData(fileContent, 1)
	result1 := day15.Run(g, end)
	fmt.Println("Part1 result : ", result1)
	g2, end2 := day15.ParseData(fileContent, 5)
	result2 := day15.Run(g2, end2)
	fmt.Println("Part2 result : ", result2)
}
