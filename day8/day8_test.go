package day8

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	messages := ParseData(fileContent)
	t.Log(messages[0])
	t.Log(messages[1])
	assert.Equal(t, 11, len(messages))
	assert.Equal(t, "dcaebfg", messages[10].input[2])
	// assert.True(t, false)
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	messages := ParseData(fileContent)
	count := Step1_CountOccurrenceOfSpecialDigitsInOutput(messages)
	assert.Equal(t, 26, count)
	// assert.True(t, false)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	messages := ParseData(fileContent)
	count := Step2_GessDigits(messages)
	assert.Equal(t, 61229+5353, count)
	// assert.True(t, false)
}

func TestPart2_guess(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	messages := ParseData(fileContent)
	guess := guess(messages[0])
	// displayRemaining(guess)
	Digit1 := readDigit(messages[0].output[0], guess)
	Digit2 := readDigit(messages[0].output[1], guess)
	Digit3 := readDigit(messages[0].output[2], guess)
	Digit4 := readDigit(messages[0].output[3], guess)
	assert.Equal(t, "5", Digit1)
	assert.Equal(t, "3", Digit2)
	assert.Equal(t, "5", Digit3)
	assert.Equal(t, "3", Digit4)
}
