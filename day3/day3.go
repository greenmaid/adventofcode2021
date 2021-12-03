package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Reading files requires checking most calls for errors. This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []string {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}

func main() {
	fmt.Println("** Days 3 **")
	// filePath := "day3/day3_input.test.txt"
	filePath := "day3/day3_input.txt"
	fileContent := readFile(filePath)
	fmt.Println("Part1 result : ", step1_calculateGamma(fileContent))
	// fmt.Println("Part2 result : ", step2_followInstructionsWithAim(fileContent))
}

func parseLineAsBits(line string) []int {
	var bits []int
	for _, bitStr := range line {
		bits = append(bits, convertRuneToInt(bitStr))
	}
	return bits
}

// https://stackoverflow.com/questions/21322173/convert-rune-to-int
func convertRuneToInt(rune rune) int {
	return int(rune - '0')
}

func sumBits(input []string) []int {

	inputLength := len(input[0])
	result := make([]int, inputLength)

	for _, line := range input {
		bits := parseLineAsBits(line)
		for idx, bit := range bits {
			result[idx] += bit
		}
	}
	return result
}

func getBits(input []string) ([]int, []int) {
	// get the sum of each digits
	sums := sumBits(input)
	fmt.Println(sums)

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
	return gammaBits, epsilonBits
}

func step1_calculateGamma(input []string) string {
	gammaBits, epsilonBits := getBits(input)
	gammaStr := ""
	epsilonStr := ""
	for _, bit := range gammaBits {
		gammaStr = fmt.Sprint(gammaStr, bit)
	}
	for _, bit := range epsilonBits {
		epsilonStr = fmt.Sprint(epsilonStr, bit)
	}

	gammaInt, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilonStr, 2, 64)

	return fmt.Sprintf("Gamma:%s / %d ;  Epsilon: %s / %d ; result %d", gammaStr, gammaInt, epsilonStr, epsilonInt, gammaInt*epsilonInt)
}
