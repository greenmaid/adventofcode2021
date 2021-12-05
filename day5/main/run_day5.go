package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day5"
	"fmt"
)

func main() {
	fmt.Println("** Day 5 **")
	filePath := "day5/input.txt"
	fileContent := common.ReadFile(filePath)
	vents := day5.ParseInputAsVentCoordinates(fileContent)
	fmt.Println("Part1 result : ", day5.Step1_calculateOverlapingAreas(vents))
	fmt.Println("Part2 result : ", day5.Step2_calculateOverlapingAreasWithDiag(vents))
}
