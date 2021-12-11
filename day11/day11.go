package day11

import (
	"adventofcode2021/common"
	"fmt"
)

type Octopus struct {
	power   int
	flashed bool
}

func ParseDataToOctopusMap(lines []string) [10][10]Octopus {

	octopusMap := [10][10]Octopus{}

	for y, line := range lines {
		powerLevels := common.ParseLineAsBits(line)
		for x, power := range powerLevels {
			octopus := Octopus{power: power, flashed: false}
			octopusMap[x][y] = octopus
		}
	}
	return octopusMap
}

func Step1_getFlashCountAfter100Steps(octopusMap [10][10]Octopus) int {
	flashCount := 0
	for step := 1; step <= 100; step++ {
		count := 0
		octopusMap, count = runStep(octopusMap)
		flashCount += count
	}
	return flashCount
}

func Step2_FindSynchronizationStep(octopusMap [10][10]Octopus) int {
	step := 0
	for {
		step += 1
		count := 0
		octopusMap, count = runStep(octopusMap)
		if count == 100 {
			break
		}
	}
	return step
}

func runStep(octopusMap [10][10]Octopus) ([10][10]Octopus, int) {

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			octopusMap = increasePower(octopusMap, x, y)
		}
	}
	NewOctopusMap, count := resetPowerForFlashed(octopusMap)

	return NewOctopusMap, count
}

func increasePower(octopusMap [10][10]Octopus, x int, y int) [10][10]Octopus {
	// octopus := octopusMap[x][y]
	octopusMap[x][y].power += 1
	if octopusMap[x][y].power > 9 && !octopusMap[x][y].flashed {
		octopusMap[x][y].flashed = true
		// octopusMap[x][y] = octopus
		for _, coord := range getNeighbors(x, y) {
			octopusMap = increasePower(octopusMap, coord[0], coord[1])
		}
	}
	return octopusMap
}

func getNeighbors(x int, y int) [][2]int {
	neighbors := make([][2]int, 0)
	if x > 0 {
		neighbors = append(neighbors, [2]int{x - 1, y})
		if y > 0 {
			neighbors = append(neighbors, [2]int{x - 1, y - 1})
		}
		if y < 9 {
			neighbors = append(neighbors, [2]int{x - 1, y + 1})
		}
	}
	if x < 9 {
		neighbors = append(neighbors, [2]int{x + 1, y})
		if y > 0 {
			neighbors = append(neighbors, [2]int{x + 1, y - 1})
		}
		if y < 9 {
			neighbors = append(neighbors, [2]int{x + 1, y + 1})
		}
	}
	if y > 0 {
		neighbors = append(neighbors, [2]int{x, y - 1})
	}
	if y < 9 {
		neighbors = append(neighbors, [2]int{x, y + 1})
	}
	return neighbors
}

func resetPowerForFlashed(octopusMap [10][10]Octopus) ([10][10]Octopus, int) {
	count := 0
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if octopusMap[x][y].flashed {
				count += 1
				octopusMap[x][y].power = 0
				octopusMap[x][y].flashed = false
			}
		}
	}
	return octopusMap, count
}

func displayGrid(octopusMap [10][10]Octopus) {

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Print(octopusMap[x][y].power)
		}
		fmt.Println("")
	}
}
