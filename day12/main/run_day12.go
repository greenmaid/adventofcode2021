package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day12"
	"fmt"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day12")
	fmt.Println("** Day 12 **")
	filePath := "day12/input.txt"
	fileContent := common.ReadFile(filePath)
	caves := day12.ParseDataToCaves(fileContent)
	count1 := day12.Step1_CountValidPaths(caves)
	fmt.Println("Part1 result : ", count1)
	count2 := day12.Step2_CountValidPaths(caves)
	fmt.Println("Part1 result : ", count2)

}
