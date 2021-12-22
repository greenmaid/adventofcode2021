package day00

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	data := ParseData(fileContent)
	assert.Equal(t, 0, len(data))
	assert.True(t, true)
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	data := ParseData(fileContent)
	result := Step1(data)
	assert.Equal(t, 1, result)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	data := ParseData(fileContent)
	result := Step2(data)
	assert.Equal(t, 2, result)
}
