package day7

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	positions := ParseDataAsPositions(fileContent)
	t.Log(positions)
	assert.Equal(t, 10, len(positions))
	assert.Equal(t, 16, positions[0])
	assert.Equal(t, 7, positions[6])
	// assert.True(t, false)
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	positions := ParseDataAsPositions(fileContent)
	minConsumption := Step1_CalculateMinFuelConsuptionForAlignment(positions)
	assert.Equal(t, 37, minConsumption)
	// assert.True(t, false)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	positions := ParseDataAsPositions(fileContent)
	minConsumption := Step2_CalculateMinFuelConsuptionForAlignmentWithIncresingCost(positions)
	assert.Equal(t, 168, minConsumption)
	// assert.True(t, false)
}
