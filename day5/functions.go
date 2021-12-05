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

// instead of using a 2d array representing map
// using a hashmap collecting all point containing some vents
// I choose to represent point x/y by the key x + 1000*y
// (I assume x is always under 1000 so there is orthogonality)
// The associate value will be the number of overlaping vent on this point

type gridmap map[int]int

func getNewGrid() gridmap {
	grid := make(gridmap)
	return grid
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
	mapGrid := getNewGrid()
	for _, vent := range vents {
		mapGrid = drawVentOnMap(mapGrid, vent, false)
	}
	overlaps := countOverlapOnMap(mapGrid)
	return fmt.Sprintf("%d", overlaps)
}

func Step2_calculateOverlapingAreasWithDiag(vents []vent) string {
	mapGrid := getNewGrid()
	for _, vent := range vents {
		mapGrid = drawVentOnMap(mapGrid, vent, true)
	}
	overlaps := countOverlapOnMap(mapGrid)
	return fmt.Sprintf("%d", overlaps)
}

func drawVentOnMap(mapGrid gridmap, vent vent, withDiagonals bool) gridmap {

	direction := 1
	if vent.start[1] == vent.end[1] {
		// horizontal vent
		if vent.start[0] > vent.end[0] {
			direction = -1
		}
		for i := 0; i <= (vent.end[0]-vent.start[0])*direction; i++ {
			mapGrid[vent.start[0]+(i*direction)+(1000*vent.start[1])] += 1
		}

	} else if vent.start[0] == vent.end[0] {
		// vertical vent
		if vent.start[1] > vent.end[1] {
			direction = -1
		}
		for i := 0; i <= (vent.end[1]-vent.start[1])*direction; i++ {
			mapGrid[vent.start[0]+1000*(vent.start[1]+(i*direction))] += 1
		}
	} else if withDiagonals {
		if vent.start[0] > vent.end[0] {
			tmpStart := vent.start
			vent.start = vent.end
			vent.end = tmpStart
		}
		if vent.start[1] > vent.end[1] {
			direction = -1
		}
		for i := 0; i <= vent.end[0]-vent.start[0]; i++ {
			mapGrid[vent.start[0]+i+1000*(vent.start[1]+(i*direction))] += 1
		}

	}
	return mapGrid
}

func countOverlapOnMap(mapGrid gridmap) int {
	count := 0
	for _, value := range mapGrid {
		if value > 1 {
			count += 1
		}
	}
	return count
}

func displayGrid(mapGrid gridmap) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Print(mapGrid[x+1000*y], " ")
		}
		fmt.Print("\n")
	}
}
