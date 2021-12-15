package day6

import (
	"strconv"
	"strings"
)

type Pool map[int]int

func calculateNextDay(pool Pool) Pool {
	nextGen := pool[0]
	for i := 0; i < 8; i++ {
		pool[i] = pool[i+1]
	}
	pool[6] += nextGen
	pool[8] = nextGen

	return pool
}

func countFishes(pool Pool) int {
	totalCount := 0
	for _, genCount := range pool {
		totalCount += genCount
	}
	return totalCount
}

func ParseDataAsLanterfishPool(fileContent []string) Pool {
	pool := Pool{}
	for _, line := range fileContent {
		for _, valueStr := range strings.Split(line, ",") {
			parsedValue, _ := strconv.Atoi(valueStr)
			pool[parsedValue] += 1
		}
	}
	return pool
}

func CounFishAfterXDays(pool Pool, days int) int {
	for i := 1; i <= days; i++ {
		pool = calculateNextDay(pool)
	}
	return countFishes(pool)
}
