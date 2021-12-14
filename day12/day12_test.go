package day12

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test1.txt"
	fileContent := common.ReadFile(filePath)
	caves := ParseDataToCaves(fileContent)
	t.Log(caves)
	assert.Equal(t, 6, len(caves))
	assert.Contains(t, caves["A"].links, "end")
	assert.NotContains(t, caves["A"].links, "d")
}

func TestPart1_1(t *testing.T) {
	filePath := "input.test1.txt"
	fileContent := common.ReadFile(filePath)
	caves := ParseDataToCaves(fileContent)
	count := Step1_CountValidPaths(caves)
	assert.Equal(t, 10, count)
}

func TestPart1_2(t *testing.T) {
	filePath := "input.test2.txt"
	fileContent := common.ReadFile(filePath)
	caves := ParseDataToCaves(fileContent)
	count := Step1_CountValidPaths(caves)
	assert.Equal(t, 19, count)
}

func TestPart1_3(t *testing.T) {
	filePath := "input.test3.txt"
	fileContent := common.ReadFile(filePath)
	caves := ParseDataToCaves(fileContent)
	count := Step1_CountValidPaths(caves)
	assert.Equal(t, 226, count)
}





func TestPart2_1(t *testing.T) {
	filePath := "input.test1.txt"
	fileContent := common.ReadFile(filePath)
	caves := ParseDataToCaves(fileContent)
	count := Step2_CountValidPaths(caves)
	assert.Equal(t, 36, count)
}

func TestPart2_2(t *testing.T) {
	filePath := "input.test2.txt"
	fileContent := common.ReadFile(filePath)
	caves := ParseDataToCaves(fileContent)
	count := Step2_CountValidPaths(caves)
	assert.Equal(t, 103, count)
}

func TestPart2_3(t *testing.T) {
	filePath := "input.test3.txt"
	fileContent := common.ReadFile(filePath)
	caves := ParseDataToCaves(fileContent)
	count := Step2_CountValidPaths(caves)
	assert.Equal(t, 3509, count)
}
