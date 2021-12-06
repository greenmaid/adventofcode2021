package day6

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	pool := ParseDataAsLanterfishPool(fileContent)
	t.Log(pool)
	assert.Equal(t, 5, countFishes(pool))
	assert.Equal(t, 0, pool[0])
	assert.Equal(t, 2, pool[3])
	// assert.True(t, false)
}

func TestPart1_FirstGenerations(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	pool := ParseDataAsLanterfishPool(fileContent)
	pool = calculateNextDay(pool)
	t.Log(pool)
	assert.Equal(t, 5, countFishes(pool))
	pool = calculateNextDay(pool)
	t.Log(pool)
	assert.Equal(t, 6, countFishes(pool))
	assert.Equal(t, 1, pool[2])
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	pool := ParseDataAsLanterfishPool(fileContent)
	for i := 1; i <= 80; i++ {
		pool = calculateNextDay(pool)
	}
	assert.Equal(t, 5934, countFishes(pool))
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	pool := ParseDataAsLanterfishPool(fileContent)
	for i := 1; i <= 256; i++ {
		pool = calculateNextDay(pool)
	}
	assert.Equal(t, 26984457539, countFishes(pool))
}
