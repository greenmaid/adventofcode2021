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

