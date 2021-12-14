package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day13"
	"fmt"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day13")
	fmt.Println("** Day 13 **")
	filePath := "day13/input.txt"
	fileContent := common.ReadFile(filePath)
	points, folds := day13.ParseData(fileContent)
	count1 := day13.Step1(points, folds)
	fmt.Println("Part1 result : ", count1)
	fmt.Println("Part2 result : ")
	day13.Step2(points, folds)

}
