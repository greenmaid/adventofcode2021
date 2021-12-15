package day9

import (
	"adventofcode2021/common"
	"fmt"
	"sort"
)

func ParseDataToMapGrid(lines []string) [][]int {

	var heightmap [][]int
	borderValue := 9
	borderLine := []int{}
	for i := 1; i <= len(lines[0])+2; i++ {
		borderLine = append(borderLine, borderValue)
	}
	heightmap = append(heightmap, borderLine)
	for _, line := range lines {
		lineWithBorder := append([]int{borderValue}, common.ParseLineAsBits(line)...)
		lineWithBorder = append(lineWithBorder, borderValue)
		heightmap = append(heightmap, lineWithBorder)
	}
	heightmap = append(heightmap, borderLine)
	return heightmap
}

type Point struct {
	x int
	y int
}

func Step1_FindLowPoints(heightmap [][]int) ([]Point, int) {
	lowPoints := []Point{}
	lowPointsRiskLevel := 0
	for x := 1; x < len(heightmap[0])-1; x++ {
		for y := 1; y < len(heightmap)-1; y++ {
			if heightmap[y][x] < heightmap[y+1][x] &&
				heightmap[y][x] < heightmap[y-1][x] &&
				heightmap[y][x] < heightmap[y][x-1] &&
				heightmap[y][x] < heightmap[y][x+1] {
				lowPoints = append(lowPoints, Point{x: x, y: y})
				lowPointsRiskLevel += heightmap[y][x] + 1
			}
		}
	}
	return lowPoints, lowPointsRiskLevel
}

func Step2_FindBassins(heightmap [][]int, lowPoints []Point) int {
	bassins := make([][]Point, 0)
	for _, lowPoint := range lowPoints {
		bassin := make([]Point, 0)
		previousPointsInBassin := len(bassin)
		bassin = append(bassin, lowPoint)
		pointsInBassin := len(bassin)
		for {
			if previousPointsInBassin == pointsInBassin {
				break
			}
			previousPointsInBassin = pointsInBassin
			for _, bassinPoint := range bassin {
				up := Point{x: bassinPoint.x, y: bassinPoint.y + 1}
				down := Point{x: bassinPoint.x, y: bassinPoint.y - 1}
				left := Point{x: bassinPoint.x - 1, y: bassinPoint.y}
				right := Point{x: bassinPoint.x + 1, y: bassinPoint.y}
				for _, point := range []Point{up, down, left, right} {
					if heightmap[point.y][point.x] != 9 &&
						!isPointInBassin(bassin, point) {
						bassin = append(bassin, point)
					}
				}
			}
			pointsInBassin = len(bassin)
		}
		bassins = append(bassins, bassin)
	}
	return scoreBassins(bassins)
}

func isPointInBassin(bassin []Point, point Point) bool {
	for _, pointInBassin := range bassin {
		if point.x == pointInBassin.x && point.y == pointInBassin.y {
			return true
		}
	}
	return false
}

func scoreBassins(bassins [][]Point) int {
	lengths := []int{}
	for _, bassin := range bassins {
		lengths = append(lengths, len(bassin))
	}
	sort.Ints(lengths)
	score := lengths[len(lengths)-1] * lengths[len(lengths)-2] * lengths[len(lengths)-3]
	return score

}

func DisplayGrid(grid [][]int) {
	for _, line := range grid {
		fmt.Println(line)
	}
}

func DisplayBassins(bassins [][]Point) {
	for _, bassin := range bassins {
		fmt.Println(bassin, "=> ", len(bassin))
	}
}
