package day3

import (
	"adventofcode2021/common"
	"fmt"
	"strconv"
)

func parseWholeInputAs2dArray(input []string) [][]int {
	var table [][]int
	for _, line := range input {
		table = append(table, common.ParseLineAsBits(line))
	}
	return table
}

func convertBinstrToInt(ListofBits []int) (string, int) {
	resultStr := ""
	for _, bit := range ListofBits {
		resultStr = fmt.Sprint(resultStr, bit)
	}
	result, _ := strconv.ParseInt(resultStr, 2, 64)
	return resultStr, int(result)
}

func Step1_calculateGammaAndEpsilon(input []string) string {
	// get input as int table
	table := parseWholeInputAs2dArray(input)
	// get the sum of each digits column
	sums := make([]int, len(table[0]))
	for raw := 0; raw < len(table); raw++ {
		for idx, bit := range table[raw] {
			sums[idx] += bit
		}
	}

	// then get each digit of gamma
	// if sum of all bits is more than half of total bit count, there is a majority of 1
	totalValues := len(input)
	gammaBits := make([]int, len(sums))
	epsilonBits := make([]int, len(sums))
	for idx, value := range sums {
		if value > (totalValues / 2) {
			gammaBits[idx] = 1
			epsilonBits[idx] = 0
		} else {
			gammaBits[idx] = 0
			epsilonBits[idx] = 1
		}
	}

	gammaStr, gammaInt := convertBinstrToInt(gammaBits)
	epsilonStr, epsilonInt := convertBinstrToInt(epsilonBits)

	return fmt.Sprintf("Gamma: %s/%d ;  Epsilon: %s/%d ; result %d", gammaStr, gammaInt, epsilonStr, epsilonInt, gammaInt*epsilonInt)
}

//################################################################################

func Step2_calculateLifeSupportRating(input []string) string {
	// get input as int table
	table := parseWholeInputAs2dArray(input)
	columnCount := len(table[0])

	o2table := table
	for index := 0; index < columnCount; index++ {
		o2table = filterTableBasedOnMajoritaryColumnBit(o2table, index)
	}

	co2table := table
	for index := 0; index < columnCount; index++ {
		co2table = filterTableBasedOnMinorityColumnBit(co2table, index)
		if len(co2table) == 1 {
			break
		}
	}

	// convert results
	o2GenRatingStr, o2GenRating := convertBinstrToInt(o2table[0])
	co2GenRatingStr, co2GenRating := convertBinstrToInt(co2table[0])

	return fmt.Sprintf("O2: %s/%d; CO2: %s/%d ; result %d", o2GenRatingStr, o2GenRating, co2GenRatingStr, co2GenRating, o2GenRating*co2GenRating)
}

func filterTableBasedOnMajoritaryColumnBit(table [][]int, filteringColumn int) [][]int {
	// get the majoritary bit of the selected column
	sum := 0
	for raw := 0; raw < len(table); raw++ {
		sum += table[raw][filteringColumn]
	}

	majorityBit := 1
	if sum*2 < len(table) {
		majorityBit = 0
	}

	// filter out any member of the table which doesn't have the majority bit
	var result [][]int
	for _, value := range table {
		if value[filteringColumn] == majorityBit {
			result = append(result, value)
		}
	}
	return result
}

func filterTableBasedOnMinorityColumnBit(table [][]int, filteringColumn int) [][]int {
	// get the majoritary bit of the selected column
	sum := 0
	for raw := 0; raw < len(table); raw++ {
		sum += table[raw][filteringColumn]
	}

	minorityBit := 0
	if sum*2 < len(table) {
		minorityBit = 1
	}

	// filter out any member of the table which doesn't have the majority bit
	var result [][]int
	for _, value := range table {
		if value[filteringColumn] == minorityBit {
			result = append(result, value)
		}
	}
	return result
}
