package day10

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_CheckChunk(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	chunkCheck0 := checkLine(fileContent[0])
	assert.False(t, chunkCheck0.err)
	chunkCheck2 := checkLine(fileContent[2])
	assert.True(t, chunkCheck2.err)
	// fmt.Println(chunkCheck2.err, string(chunkCheck2.errChar))
	// assert.True(t, false)
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	score := Step1_CheckData(fileContent)
	assert.Equal(t, 26397, score)
	// assert.True(t, false)
}

func TestPart2_CheckIncompleteChunk(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	chunkCheck0 := checkLine(fileContent[0])
	assert.True(t, chunkCheck0.incomplete)
	assert.Equal(t, "}}]])})]", string(chunkCheck0.finishingChars))
	// fmt.Println(chunkCheck0.incomplete, string(chunkCheck0.finishingChars))
	chunkCheck2 := checkLine(fileContent[2])
	assert.False(t, chunkCheck2.incomplete)
	// fmt.Println(chunkCheck2.incomplete, string(chunkCheck2.finishingChars))
	// assert.True(t, false)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	score := Step2_CheckIncompleteData(fileContent)
	assert.Equal(t, 288957, score)
	// assert.True(t, false)
}
