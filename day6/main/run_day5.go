package main

import (
	"adventofcode2021/common"
	"adventofcode2021/day6"
	"fmt"
	"log"
	// "time"
)

func main() {
	// defer common.TimeTrack(time.Now(), "day6")
	fmt.Println("** Day 6 **")
	filePath := "day6/input.txt"
	fileContent := common.ReadFile(filePath)
	pool := day6.ParseDataAsLanterfishPool(fileContent)
	log.Printf("Fish count: %d", len(pool))
	fmt.Println("Part1 result : ", day6.CounFishAfterXDays(pool, 80))
	pool2 := day6.ParseDataAsLanterfishPool(fileContent)
	fmt.Println("Part2 result : ", day6.CounFishAfterXDays(pool2, 256))
}
