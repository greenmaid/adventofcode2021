package day17

import (
	"fmt"
	"regexp"
	"strconv"
)

type targetArea struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func ParseData(input string) targetArea {
	re := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	match := re.FindStringSubmatch(input)
	xmin, _ := strconv.Atoi(match[1])
	xmax, _ := strconv.Atoi(match[2])
	ymin, _ := strconv.Atoi(match[3])
	ymax, _ := strconv.Atoi(match[4])
	return targetArea{
		xmin: xmin,
		xmax: xmax,
		ymin: ymin,
		ymax: ymax,
	}
}

func Step1(targetArea targetArea) int {
	maxHeight := 0
	X := 0
	Y := 0
	for y := 0; y < 200; y++ {
		maxHeightForThisInitialY := 0
		thisX := 0
		for x := 0; x <= targetArea.xmax; x++ {
			isSuccess, height, _ := isReachingTarget(targetArea, [2]int{x, y})
			if isSuccess && height > maxHeightForThisInitialY {
				maxHeightForThisInitialY = height
				thisX = x
			}
		}
		if maxHeight <= maxHeightForThisInitialY {
			maxHeight = maxHeightForThisInitialY
			X = thisX
			Y = y
		}
	}
	fmt.Printf("Result: height=%d X=%d Y=%d\n", maxHeight, X, Y)
	return maxHeight
}

func Step2(targetArea targetArea) int {
	count := 0
	for y := targetArea.ymin; y < 200; y++ {
		for x := 0; x <= targetArea.xmax; x++ {
			isSuccess, _, _ := isReachingTarget(targetArea, [2]int{x, y})
			if isSuccess {
				count++
			}
		}
	}
	return count
}

func isReachingTarget(targetArea targetArea, vector [2]int) (bool, int, int) {

	position := [2]int{0, 0}
	vectx := vector[0]
	vecty := vector[1]
	step := 0
	maxHeight := 0
	for {
		position[0] += vectx
		position[1] += vecty
		step++

		if position[1] > maxHeight {
			maxHeight = position[1]
		}

		// assert x is always >= 0
		if vectx > 0 {
			vectx--
		}
		vecty--

		if isOnTarget(position, targetArea) {
			return true, maxHeight, step
		}

		if willFail(position, targetArea, vectx, vecty) {
			return false, 0, step
		}
	}
}

func isOnTarget(position [2]int, targetArea targetArea) bool {
	x := position[0]
	y := position[1]
	if x >= targetArea.xmin &&
		x <= targetArea.xmax &&
		y >= targetArea.ymin &&
		y <= targetArea.ymax {
		return true
	}
	return false
}

// determine that shot is missed
func willFail(position [2]int, targetArea targetArea, Xspeed int, Yspeed int) bool {
	x := position[0]
	y := position[1]
	if x > targetArea.xmax {
		return true
	}
	if x < targetArea.xmin && Xspeed == 0 {
		return true
	}
	if y < targetArea.ymin {
		return true
	}
	return false
}
