package day15

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	g, end := ParseData(fileContent, 1)
	// t.Log(g)
	assert.Equal(t, 2, g.Degree(0))  // Degree return number of outward directed edges from a vertice
	assert.Equal(t, 0, g.Degree(99)) // 99 doesn't exist

	assert.Equal(t, int64(9), g.Cost(3001, 3002))
	assert.Equal(t, int64(9), g.Cost(2002, 3002))

	assert.Equal(t, 9009, end)
	assert.Equal(t, int64(1), g.Cost(end-1, end))
	// assert.True(t, false)
}

func TestParseData2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	g, end := ParseData(fileContent, 5)
	// t.Log(g)
	assert.Equal(t, 2, g.Degree(0))  // Degree return number of outward directed edges from a vertice
	assert.Equal(t, 0, g.Degree(99)) // 99 doesn't exist
	assert.Equal(t, 3, g.Degree(10))

	assert.Equal(t, int64(9), g.Cost(3001, 3002))
	assert.Equal(t, int64(9), g.Cost(2002, 3002))

	assert.Equal(t, 49049, end)
	assert.Equal(t, int64(9), g.Cost(end-1, end))
	assert.Equal(t, int64(7), g.Cost(end-2, end-1))
	// assert.True(t, false)
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	g, end := ParseData(fileContent, 1)
	result := Run(g, end)
	assert.Equal(t, int64(40), result)
	// assert.True(t, false)
}

func TestPart1_IndirectWay(t *testing.T) {
	filePath := "input.test2.txt"
	fileContent := common.ReadFile(filePath)
	g, end := ParseData(fileContent, 1)
	result := Run(g, end)
	assert.Equal(t, int64(8), result)
	// assert.True(t, false)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	g, end := ParseData(fileContent, 5)
	result := Run(g, end)
	assert.Equal(t, int64(315), result)
}
