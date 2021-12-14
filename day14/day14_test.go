package day14

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	polymer, rules := ParseData(fileContent)
	t.Log(polymer)
	t.Log(rules)
	assert.Equal(t, 2, polymer.elemCount['N'])
	assert.Equal(t, 0, polymer.elemCount['H'])
	assert.Equal(t, 1, polymer.pairs[[2]byte{'N', 'C'}])
	assert.Equal(t, byte('C'), rules[[2]byte{'H', 'N'}])
}

func TestPolymerization(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	polymer, rules := ParseData(fileContent)
	t.Log(polymer)

	polymer1 := polymerization(polymer, rules)
	t.Log(polymer1)
	assert.Equal(t, 2, polymer1.elemCount['N'])
	assert.Equal(t, 1, polymer1.elemCount['H'])
	assert.Equal(t, 2, polymer1.elemCount['C'])
	assert.Equal(t, 1, polymer.pairs[[2]byte{'N', 'C'}])

	polymer2 := polymerization(polymer1, rules)
	t.Log(polymer2)
	assert.Equal(t, 1, polymer1.elemCount['H'])
	assert.Equal(t, 4, polymer2.elemCount['C'])
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	polymer, rules := ParseData(fileContent)
	result := Step1(polymer, rules)
	assert.Equal(t, 1588, result)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	polymer, rules := ParseData(fileContent)
	result := Step2(polymer, rules)
	assert.Equal(t, 2188189693529, result)
	// assert.True(t, false)
}
