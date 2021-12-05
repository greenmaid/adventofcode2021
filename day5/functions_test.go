package day5

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	vents := ParseInputAsVentCoordinates(fileContent)
	t.Log(vents)
	assert.Equal(t, 10, len(vents))
	expected := vent{start: [2]int{0, 9}, end: [2]int{2, 9}}
	assert.Equal(t, expected, vents[6])
}

func TestVentDrawing(t *testing.T) {
	mapGrid := getNewGrid()
	vent1 := vent{start: [2]int{0, 0}, end: [2]int{0, 0}}
	mapGrid = drawVentOnMap(mapGrid, vent1, false)
	assert.Equal(t, 1, mapGrid[0])

	vent2 := vent{start: [2]int{0, 9}, end: [2]int{1, 9}}
	mapGrid = drawVentOnMap(mapGrid, vent2, false)
	assert.Equal(t, 1, mapGrid[0+9000])
	assert.Equal(t, 1, mapGrid[1+9000])

	vent3 := vent{start: [2]int{1, 1}, end: [2]int{3, 3}}
	mapGrid = drawVentOnMap(mapGrid, vent3, true)
	assert.Equal(t, 1, mapGrid[2+2000])
	displayGrid(mapGrid)

	vent4 := vent{start: [2]int{1, 3}, end: [2]int{3, 1}}
	mapGrid = drawVentOnMap(mapGrid, vent4, true)
	assert.Equal(t, 2, mapGrid[2+2000])
	displayGrid(mapGrid)

	vent5 := vent{start: [2]int{4, 4}, end: [2]int{2, 6}}
	mapGrid = drawVentOnMap(mapGrid, vent5, true)
	assert.Equal(t, 1, mapGrid[3+5000])
	displayGrid(mapGrid)

	vent6 := vent{start: [2]int{4, 6}, end: [2]int{2, 4}}
	mapGrid = drawVentOnMap(mapGrid, vent6, true)
	assert.Equal(t, 2, mapGrid[3+5000])
	displayGrid(mapGrid)
}

func Test_part1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	vents := ParseInputAsVentCoordinates(fileContent)

	count := Step1_calculateOverlapingAreas(vents)
	assert.Equal(t, "5", count)
}

func Test_part2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	vents := ParseInputAsVentCoordinates(fileContent)
	count := Step2_calculateOverlapingAreasWithDiag(vents)
	assert.Equal(t, "12", count)
}
