package day18

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseData(input []string) []string {
	return input
}

func Step1(snailNumbers []string) int {
	snailNumber := snailNumbers[0]
	for i := 1; i < len(snailNumbers); i++ {
		snailNumber = addAndReducePair(snailNumber, snailNumbers[i])
	}
	return calculateMagnitude(snailNumber)
}

func Step2(snailNumbers []string) int {
	maxMagnitude := 0
	for one := 0; one < len(snailNumbers); one++ {
		for two := 0; two < len(snailNumbers); two++ {
			num := addAndReducePair(snailNumbers[one], snailNumbers[two])
			magnitude := calculateMagnitude(num)
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}
	return maxMagnitude
}

func getPair(snailNumber string) [2]string {
	separatorIndex := -1
	depth := 0
	for index, char := range snailNumber {
		if char == '[' {
			depth += 1
		}
		if char == ']' {
			depth -= 1
		}
		if char == ',' && depth == 1 {
			separatorIndex = index
			break
		}
	}
	if separatorIndex == -1 {
		fmt.Println("ERROR: unable to find pair in ", snailNumber)
		panic("Pair not found")
	}
	left := snailNumber[1:separatorIndex]
	right := snailNumber[separatorIndex+1 : len(snailNumber)-1]
	return [2]string{left, right}
}

func explodePair(snailNumber string, depth int) (string, int, int) {
	pair := getPair(snailNumber)

	pair0Explosed := pair[0]
	pair1Explosed := pair[1]
	pair0LeftReminder := 0
	pair0RightReminder := 0
	pair1LeftReminder := 0
	pair1RightReminder := 0

	if depth < 4 {

		if strings.Contains(pair[0], "[") {
			pair0Explosed, pair0LeftReminder, pair0RightReminder = explodePair(pair[0], depth+1)

		}
		if pair0Explosed == pair[0] && strings.Contains(pair[1], "[") {
			pair1Explosed, pair1LeftReminder, pair1RightReminder = explodePair(pair[1], depth+1)
		}
		if pair0RightReminder > 0 {
			pair1Explosed = addRight(pair1Explosed, pair0RightReminder)
		}
		if pair1LeftReminder > 0 {
			pair0Explosed = addLeft(pair0Explosed, pair1LeftReminder)
		}
		// fmt.Println(pair0RightReminder, pair1LeftReminder)
		return fmt.Sprintf("[%s,%s]", pair0Explosed, pair1Explosed), pair0LeftReminder, pair1RightReminder
	}

	pair0, _ := strconv.Atoi(pair[0])
	pair1, _ := strconv.Atoi(pair[1])
	return "0", pair0, pair1

}

func addLeft(snailNumber string, addValue int) string {
	re := regexp.MustCompile(`(\d+)(]*)$`)
	match := re.FindStringSubmatch(snailNumber)
	initialValue, _ := strconv.Atoi(match[1])
	return re.ReplaceAllString(snailNumber, fmt.Sprint(initialValue+addValue, match[2]))
}

func addRight(snailNumber string, addValue int) string {
	re := regexp.MustCompile(`^(\[*)(\d+)`)
	match := re.FindStringSubmatch(snailNumber)
	initialValue, _ := strconv.Atoi(match[2])
	return re.ReplaceAllString(snailNumber, fmt.Sprint(match[1], initialValue+addValue))
}

func splitPair(snailNumber string) string {

	findNumberRe := regexp.MustCompile(`(\d+)`)
	match := findNumberRe.FindAllString(snailNumber, -1)

	for _, value := range match {
		numVal, _ := strconv.Atoi(value)
		if numVal >= 10 {
			newPair := fmt.Sprintf("[%d,%d]", numVal/2, numVal/2+numVal%2)

			re := regexp.MustCompile(fmt.Sprintf("^(.*?)%d(.*)$", numVal))
			repl := fmt.Sprintf("${1}%s$2", newPair)
			return re.ReplaceAllString(snailNumber, repl)

		}
	}
	return snailNumber
}

func addAndReducePair(snailNumber1 string, snailNumber2 string) string {
	newNumber := fmt.Sprintf("[%s,%s]", snailNumber1, snailNumber2)
	previousNumber := ""

	for {
		if newNumber == previousNumber {
			break
		}
		previousNumber = newNumber
		newNumber, _, _ = explodePair(newNumber, 0)
		if newNumber != previousNumber {
			continue
		}
		newNumber = splitPair(newNumber)
	}
	return newNumber
}

func addSnailNumbers(snailNumbers []string) string {
	snailNumber := snailNumbers[0]
	for i := 1; i < len(snailNumbers); i++ {
		// fmt.Println("+ ", snailNumbers[i])
		snailNumber = addAndReducePair(snailNumber, snailNumbers[i])
		// fmt.Println("= ", snailNumber)
	}

	return snailNumber
}

func calculateMagnitude(snailNumber string) int {
	pair := getPair(snailNumber)
	leftMagnitude := 0
	rightMagnitude := 0

	if strings.Contains(pair[0], "[") {
		leftMagnitude = 3 * calculateMagnitude(pair[0])
	} else {
		value, _ := strconv.Atoi(pair[0])
		leftMagnitude = 3 * value
	}

	if strings.Contains(pair[1], "[") {
		rightMagnitude = 2 * calculateMagnitude(pair[1])
	} else {
		value, _ := strconv.Atoi(pair[1])
		rightMagnitude = 2 * value
	}
	return leftMagnitude + rightMagnitude
}
