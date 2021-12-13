package day13

import (
	"fmt"
	"regexp"
	"strconv"
)

type folding struct {
	axe   string
	value int
}

func ParseData(lines []string) (map[[2]int]bool, []folding) {
	pointRegExp := regexp.MustCompile(`(\d+),(\d+)`)
	foldRegExp := regexp.MustCompile(`fold along (x|y)=(\d+)`)
	points := make(map[[2]int]bool)
	folds := make([]folding, 0)
	for _, line := range lines {
		matchPoint := pointRegExp.FindStringSubmatch(line)
		if len(matchPoint) > 0 {
			x, _ := strconv.Atoi(matchPoint[1])
			y, _ := strconv.Atoi(matchPoint[2])
			points[[2]int{x, y}] = true
		}

		matchFold := foldRegExp.FindStringSubmatch(line)
		if len(matchFold) > 0 {
			axe := matchFold[1]
			value, _ := strconv.Atoi(matchFold[2])
			folds = append(folds, folding{axe: axe, value: value})
		}

	}
	return points, folds
}

func Step1(points map[[2]int]bool, folds []folding) int {
	newPoints := foldPaper(points, folds[0])
	return len(newPoints)

}

func Step2(points map[[2]int]bool, folds []folding) {

	currentPoints := points
	for _, folding := range folds {
		newPoints := foldPaper(currentPoints, folding)
		currentPoints = newPoints
	}

	displayPoints(currentPoints)

}

func foldPaper(points map[[2]int]bool, fold folding) map[[2]int]bool {
	newPoints := make(map[[2]int]bool)
	for points := range points {
		X := points[0]
		Y := points[1]
		if fold.axe == "x" {
			if X > fold.value {
				newPoints[[2]int{2*fold.value - X, Y}] = true
			} else {
				newPoints[[2]int{X, Y}] = true
			}
		} else if fold.axe == "y" {
			if Y > fold.value {
				newPoints[[2]int{X, 2*fold.value - Y}] = true
			} else {
				newPoints[[2]int{X, Y}] = true
			}
		}
	}
	return newPoints
}

func displayPoints(points map[[2]int]bool) {
	minX := 10000
	minY := 10000
	maxX := 0
	maxY := 0
	for point := range points {
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[0] < minX {
			minX = point[0]
		}
		if point[1] > maxY {
			maxY = point[0]
		}
		if point[1] < minY {
			minY = point[0]
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := points[[2]int{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
