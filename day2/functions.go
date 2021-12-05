package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type submarine struct {
	position int
	depth    int
}

func Step1_followInstructions(instructionsList []string) string {

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

func Step2_followInstructionsWithAim(instructionsList []string) string {

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
