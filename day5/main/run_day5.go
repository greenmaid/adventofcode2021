package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day5"
	"fmt"
	"log"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day5")
	fmt.Println("** Day 5 **")
	filePath := "day5/input.txt"
	fileContent := common.ReadFile(filePath)
	vents := day5.ParseInputAsVentCoordinates(fileContent)
	log.Printf("Vent count: %d", len(vents))
	fmt.Println("Part1 result : ", day5.Step1_calculateOverlapingAreas(vents))
	fmt.Println("Part2 result : ", day5.Step2_calculateOverlapingAreasWithDiag(vents))
}
