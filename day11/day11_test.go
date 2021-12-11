package day11

import (
	"adventofcode2021/common"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	OctopusMap := ParseDataToOctopusMap(fileContent)
	t.Log(OctopusMap)
	assert.Equal(t, 5, OctopusMap[0][0].power)
	assert.Equal(t, 1, OctopusMap[3][3].power)
}

func TestPar1_FistStep(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	octopusMap := ParseDataToOctopusMap(fileContent)
	flashCount := 0
	octopusMap, flashCount = runStep(octopusMap)
	displayGrid(octopusMap)
	assert.Equal(t, 4, octopusMap[4][4].power)
	assert.Equal(t, 0, flashCount)
	octopusMap, flashCount = runStep(octopusMap)
	fmt.Println("")
	displayGrid(octopusMap)
	assert.Equal(t, 0, octopusMap[0][7].power)
	assert.Equal(t, 0, octopusMap[8][4].power)
	assert.Equal(t, 35, flashCount)
}

func TestPar1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	octopusMap := ParseDataToOctopusMap(fileContent)
	flashCount := Step1_getFlashCountAfter100Steps(octopusMap)
	assert.Equal(t, 1656, flashCount)

}

func TestPar2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	octopusMap := ParseDataToOctopusMap(fileContent)
	count := Step2_FindSynchronizationStep(octopusMap)
	assert.Equal(t, 195, count)

}
