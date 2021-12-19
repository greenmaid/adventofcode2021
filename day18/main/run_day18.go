package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day18"
	"fmt"
)

func main() {
	fmt.Println("** Day 18 **")
	filePath := "day18/input.txt"
	data := common.ReadFile(filePath)
	result1 := day18.Step1(data)
	fmt.Println("Part1 result : ", result1)
	result2 := day18.Step2(data)
	fmt.Println("Part2 result : ", result2)
}
