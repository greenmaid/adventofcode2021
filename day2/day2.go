package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("** Days 2 **")
	// filePath := "day2/day2_input.test.txt"
	filePath := "day2/day2_input.txt"
	fileContent := readFile(filePath)
	fmt.Println("Part1 result : ", step1_followInstructions(fileContent))
	fmt.Println("Part2 result : ", step2_followInstructionsWithAim(fileContent))
}

type submarine struct {
	position int
	depth    int
}

func step1_followInstructions(instructionsList []string) string {

	sub := submarine{0, 0}

	for _, instruction := range instructionsList {
		parsedAction := strings.Fields(instruction)
		action := parsedAction[0]
		actionValue, _ := strconv.Atoi(parsedAction[1])
		switch action {
		case "forward":
			sub.position += actionValue
		case "up":
			sub.depth -= actionValue
		case "down":
			sub.depth += actionValue
		default:
			log.Fatalf("Unknown action %s", action)
		}

	}

	return fmt.Sprintf("position => %d ; Depth => %d ; result => %d", sub.position, sub.depth, sub.position*sub.depth)
}

type submarine2 struct {
	aim      int
	position int
	depth    int
}

func step2_followInstructionsWithAim(instructionsList []string) string {

	sub := submarine2{0, 0, 0}

	for _, instruction := range instructionsList {
		parsedAction := strings.Fields(instruction)
		action := parsedAction[0]
		actionValue, _ := strconv.Atoi(parsedAction[1])
		switch action {
		case "forward":
			sub.position += actionValue
			sub.depth += sub.aim * actionValue
		case "up":
			sub.aim -= actionValue
		case "down":
			sub.aim += actionValue
		default:
			log.Fatalf("Unknown action %s", action)
		}

	}

	return fmt.Sprintf("position => %d ; Depth => %d ; result => %d", sub.position, sub.depth, sub.position*sub.depth)
}
