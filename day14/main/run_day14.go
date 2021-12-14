package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day14"
	"fmt"
	"time"
	// "time"
)

func main() {
	fmt.Println("** Day 14 **")
	filePath := "day14/input.txt"
	fileContent := common.ReadFile(filePath)
	polymer, rules := day14.ParseData(fileContent)
	count1 := day14.Step1(polymer, rules)
	fmt.Println("Part1 result : ", count1)
	polymer2, rules2 := day14.ParseData(fileContent)
	defer common.TimeTrack(time.Now(), "day14")
	count2 := day14.Step2(polymer2, rules2)
	fmt.Println("Part2 result : ", count2)

}
