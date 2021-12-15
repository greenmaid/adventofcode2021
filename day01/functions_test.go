package day1

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1_1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFileToInt(filePath)
	assert.Equal(t, 7, Step1_countIncreaseValues(fileContent))
}

func TestDay1_2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFileToInt(filePath)
	assert.Equal(t, 5, Step2_countIncreaseValuesByGroup(fileContent))
}
