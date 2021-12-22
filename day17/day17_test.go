package day17

import (
	"adventofcode2021/common"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	target := ParseData(fileContent[0])
	assert.Equal(t, 20, target.xmin)
	assert.Equal(t, 30, target.xmax)
	assert.Equal(t, -10, target.ymin)
	assert.Equal(t, -5, target.ymax)
}

func TestPart1(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	target := ParseData(fileContent[0])
	result := Step1(target)
	assert.Equal(t, 45, result)
}

func TestPart2(t *testing.T) {
	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	target := ParseData(fileContent[0])
	result := Step2(target)
	assert.Equal(t, 112, result)
}

func Test_isReachingTarget(t *testing.T) {

	filePath := "input.test.txt"
	fileContent := common.ReadFile(filePath)
	target := ParseData(fileContent[0])

	fmt.Println(target)

	tests := []struct {
		name   string
		vector [2]int
		want   bool
		height int
	}{
		{name: "test1", vector: [2]int{0, 0}, want: false, height: 0},
		{name: "test2", vector: [2]int{7, 2}, want: true, height: 3},
		{name: "test3", vector: [2]int{6, 3}, want: true, height: 6},
		{name: "test4", vector: [2]int{9, 0}, want: true, height: 0},
		{name: "test5", vector: [2]int{17, -4}, want: false, height: 0},
		{name: "test6", vector: [2]int{6, 9}, want: true, height: 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSuccess, height, _ := isReachingTarget(target, tt.vector)
			assert.Equalf(t, tt.want, isSuccess, "%s result failed => %v, want %v", tt.name, isSuccess, tt.want)
			assert.Equalf(t, tt.height, height, "%s height failed => %v, want %v", tt.name, height, tt.height)
		})

	}
}
