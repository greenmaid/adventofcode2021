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

func readFileToInt(path string) []int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	var content []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scannedText, _ := strconv.Atoi(scanner.Text())
		content = append(content, scannedText)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}

func main() {
	fmt.Println("** Days 1 **")
	// filePath := "day1/day1_input.test.txt"
	filePath := "day1/day1_input.txt"
	fileContent := readFileToInt(filePath)
	fmt.Println("Part1 result : ", step1_countIncreaseValues(fileContent))
	fmt.Println("Part2 result : ", step2_countIncreaseValuesByGroup(fileContent))
}

func step1_countIncreaseValues(values []int) int {
	count := 0
	for index := range values {
		if index > 0 {
			if values[index-1] < values[index] {
				count += 1
			}
		}
	}
	return count
}

func step2_countIncreaseValuesByGroup(values []int) int {

	var groupedValues []int
	for index := range values {
		if index < len(values)-2 {
			groupedValues = append(groupedValues, values[index]+values[index+1]+values[index+2])
		}
	}
	return step1_countIncreaseValues(groupedValues)
}
