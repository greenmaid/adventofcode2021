package day7

import (
	"strconv"
	"strings"
)

func ParseDataAsPositions(fileContent []string) []int {
	result := make([]int, 0)
	for _, line := range fileContent {
		for _, valueStr := range strings.Split(line, ",") {
			parsedValue, _ := strconv.Atoi(valueStr)
			result = append(result, parsedValue)
		}
	}
	return result
}

func Step1_CalculateMinFuelConsuptionForAlignment(positions []int) int {
	min := getMinFromList(positions)
	max := getMaxFromList(positions)
	minFuelConsumption := calculateFuelConsumption(positions, 0)
	for i := min; i <= max; i++ {
		minFuelConsumption = minimun(minFuelConsumption, calculateFuelConsumption(positions, i))
	}
	return minFuelConsumption
}

func Step2_CalculateMinFuelConsuptionForAlignmentWithIncresingCost(positions []int) int {
	min := getMinFromList(positions)
	max := getMaxFromList(positions)
	minFuelConsumption := calculateFuelConsumptionWithIncreasingCost(positions, 0)
	for i := min; i <= max; i++ {
		minFuelConsumption = minimun(minFuelConsumption, calculateFuelConsumptionWithIncreasingCost(positions, i))
	}
	return minFuelConsumption
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimun(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

func calculateFuelConsumption(positions []int, alignmentPosition int) int {
	fuel := 0
	for _, currentPosition := range positions {
		fuel += abs(currentPosition - alignmentPosition)
	}
	return fuel
}

func calculateFuelConsumptionWithIncreasingCost(positions []int, alignmentPosition int) int {
	fuel := 0
	for _, currentPosition := range positions {
		dist := abs(currentPosition - alignmentPosition)
		cost := dist * (dist + 1) / 2
		fuel += cost
	}
	return fuel
}

func getMinFromList(list []int) int {
	min := 1000 // arbitrary, works here but cheating
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	return min
}

func getMaxFromList(list []int) int {
	max := 0
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}
