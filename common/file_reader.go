package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileToInt(path string) []int {
	file, err := os.Open(path)
	Check(err)
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

func ReadFile(path string) []string {
	file, err := os.Open(path)
	Check(err)
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

// get each line as a list of integer
func ParseLineAsBits(line string) []int {
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
