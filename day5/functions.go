package day5

import (
	"fmt"
	"regexp"
	"strconv"
)

type vent struct {
	start [2]int
	end   [2]int
}

func ParseInputAsVentCoordinates(input []string) []vent {

	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	vents := []vent{}
	for _, line := range input {
		match := re.FindStringSubmatch(line)
		startX, _ := strconv.Atoi(match[1])
		startY, _ := strconv.Atoi(match[2])
		endX, _ := strconv.Atoi(match[3])
		endY, _ := strconv.Atoi(match[4])
		vents = append(vents, vent{start: [2]int{startX, startY}, end: [2]int{endX, endY}})
	}
	return vents
}

func Step1_calculateOverlapingAreas(vents []vent) string {
	mapGrid := [1000][1000]int{}
	for _, vent := range vents {
		mapGrid = drawVentOnMap(mapGrid, vent, false)
	}
	overlaps := countOverlapOnMap(mapGrid)
	return fmt.Sprintf("%d", overlaps)
}

func Step2_calculateOverlapingAreasWithDiag(vents []vent) string {
	mapGrid := [1000][1000]int{}
	for _, vent := range vents {
		mapGrid = drawVentOnMap(mapGrid, vent, true)
	}
	overlaps := countOverlapOnMap(mapGrid)
	return fmt.Sprintf("%d", overlaps)
}

func drawVentOnMap(mapGrid [1000][1000]int, vent vent, withDiagonals bool) [1000][1000]int {

	orientedVent := vent
	if vent.start[0] > vent.end[0] {
		orientedVent.start = vent.end
		orientedVent.end = vent.start
	}

	if orientedVent.start[1] == orientedVent.end[1] {
		// horizontal vent
		for x := orientedVent.start[0]; x <= orientedVent.end[0]; x++ {
			mapGrid[x][orientedVent.start[1]] += 1
		}
	} else if orientedVent.start[0] == orientedVent.end[0] {
		// vertical vent
		if orientedVent.start[1] > orientedVent.end[1] {
			orientedVent.start = vent.end
			orientedVent.end = vent.start
		}
		for y := orientedVent.start[1]; y <= orientedVent.end[1]; y++ {
			mapGrid[orientedVent.start[0]][y] += 1
		}
	} else if withDiagonals {
		direction := 1
		if orientedVent.start[1] > orientedVent.end[1] {
			direction = -1
		}
		for i := 0; i <= orientedVent.end[0]-orientedVent.start[0]; i++ {
			mapGrid[orientedVent.start[0]+i][orientedVent.start[1]+(i*direction)] += 1
		}

	}
	return mapGrid
}

func countOverlapOnMap(mapGrid [1000][1000]int) int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if mapGrid[x][y] > 1 {
				count += 1
			}
		}
	}
	return count
}

func displayGrid(mapGrid [1000][1000]int) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Print(mapGrid[x][y], " ")
		}
		fmt.Print("\n")
	}
}
