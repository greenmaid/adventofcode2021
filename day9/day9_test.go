package day9

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	grid := ParseDataToMapGrid(fileContent)
	// DisplayGrid(grid)
	assert.Equal(t, 5+2, len(grid))
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	heightmap := ParseDataToMapGrid(fileContent)
	_, count := Step1_FindLowPoints(heightmap)
	assert.Equal(t, 15, count)
	// DisplayGrid(heightmap)
	// assert.True(t, false)

}
func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	heightmap := ParseDataToMapGrid(fileContent)
	lowPoints, _ := Step1_FindLowPoints(heightmap)
	count2 := Step2_FindBassins(heightmap, lowPoints)
	assert.Equal(t, 1134, count2)
	// assert.True(t, false)

}
