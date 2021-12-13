package day13

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	points, folds := ParseData(fileContent)
	// t.Log(points)
	assert.Len(t, points, 18)
	assert.Equal(t, true, points[[2]int{4, 1}])
	assert.Contains(t, points, [2]int{2, 14})
	assert.NotContains(t, points, [2]int{0, 5})
	// t.Log(folds)
	assert.Len(t, folds, 2)
	assert.Equal(t, "y", folds[0].axe)
	assert.Equal(t, 7, folds[0].value)
	assert.Equal(t, "x", folds[1].axe)
	assert.Equal(t, 5, folds[1].value)
	// assert.True(t, false)
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	points, folds := ParseData(fileContent)
	count := Step1(points, folds)
	assert.Equal(t, 17, count)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	points, folds := ParseData(fileContent)
	Step2(points, folds)
	// assert.Equal(t, 0, 1)
}
