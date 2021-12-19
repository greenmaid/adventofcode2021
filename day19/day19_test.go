package day19

import (
	"adventofcode2021/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {
	fileContent := common.ReadFile("input.test.txt")
	data := ParseData(fileContent)
	t.Log(data)
	assert.Equal(t, 5, len(data))
	assert.True(t, true)
}

func TestRun(t *testing.T) {
	fileContent := common.ReadFile("input.test.txt")
	scanners := ParseData(fileContent)
	result1, result2 := Run(scanners)
	assert.Equal(t, 79, result1)
	assert.Equal(t, 3621, result2)
}

// func TestPart2(t *testing.T) {
// 	filePath := "input.test.txt"
// 	fileContent := common.ReadFile(filePath)
// 	data := ParseData(fileContent)
// 	result := Step2(data)
// 	assert.Equal(t, 2, result)
// }

func Test_countMatchingElements(t *testing.T) {
	type args struct {
		list1 []float64
		list2 []float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{list1: []float64{1, 2, 3}, list2: []float64{1, 2, 3}}, want: 3},
		{name: "test2", args: args{list1: []float64{1, 2, 2}, list2: []float64{1, 2, 3}}, want: 2},
		{name: "test3", args: args{list1: []float64{1, 2, 2, 3}, list2: []float64{1, 2, 3, 3}}, want: 3},
		{name: "test4", args: args{list1: []float64{1, 2, 2, 3}, list2: []float64{1, 2, 3, 3}}, want: 3},
		{name: "test5", args: args{list1: []float64{1}, list2: []float64{2, 3, 3}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatchingElements(tt.args.list1, tt.args.list2); got != tt.want {
				t.Errorf("countMatchingElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPossiblematches(t *testing.T) {
	fileContent := common.ReadFile("input.test1.txt")
	scanners := ParseData(fileContent)
	tests := []struct {
		name              string
		beacon            [3]int
		scanner1          scanner
		scanner2          scanner
		possibleMatches   []int
		matchingDistCount int
	}{
		{name: "test1", beacon: scanners[0].beacons[0], scanner1: scanners[0], scanner2: scanners[1], possibleMatches: []int{1}, matchingDistCount: 3},
		{name: "test2", beacon: scanners[0].beacons[1], scanner1: scanners[0], scanner2: scanners[1], possibleMatches: []int{0}, matchingDistCount: 3},
		{name: "test3", beacon: scanners[0].beacons[2], scanner1: scanners[0], scanner2: scanners[1], possibleMatches: []int{2}, matchingDistCount: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			possibleMatches, matchingDistCount := findPossiblematches(tt.beacon, tt.scanner1, tt.scanner2)
			t.Log(possibleMatches)
			t.Log(matchingDistCount)
			assert.Equal(t, possibleMatches, tt.possibleMatches)
			assert.Equal(t, matchingDistCount, tt.matchingDistCount)
		})
	}
}

func Test_findMatches(t *testing.T) {
	fileContent := common.ReadFile("input.test.txt")
	scanners := ParseData(fileContent)
	tests := []struct {
		name     string
		scanner1 scanner
		scanner2 scanner
	}{
		{name: "test1", scanner1: scanners[0], scanner2: scanners[1]},
		{name: "test1", scanner1: scanners[1], scanner2: scanners[3]},
		{name: "test1", scanner1: scanners[1], scanner2: scanners[3]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := findMatches(tt.scanner1, tt.scanner2)
			t.Log(matches)
			assert.Equal(t, 12, len(matches))
		})
	}
}

func Test_getScannerTranslation(t *testing.T) {
	fileContent := common.ReadFile("input.test.txt")
	scanners := ParseData(fileContent)
	matches := findMatches(scanners[0], scanners[1])
	scanCoords, scanInOtherRef := getScannerTranslation(scanners[0], scanners[1], matches)
	t.Log(scanCoords)
	t.Log(scanInOtherRef)
	t.Log(len(scanInOtherRef.beacons))
	// assert.True(t, false)
}

func Test_mergeScanners(t *testing.T) {
	fileContent := common.ReadFile("input.test.txt")
	scanners := ParseData(fileContent)
	matches := findMatches(scanners[0], scanners[1])
	_, scanInOtherRef := getScannerTranslation(scanners[0], scanners[1], matches)
	generatedScanner := mergeScanners(scanners[0], scanInOtherRef)
	t.Log(generatedScanner)
	t.Log(len(generatedScanner.beacons))
	// assert.True(t, false)
}
