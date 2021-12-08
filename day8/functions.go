package day8

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type message struct {
	input  []string
	output []string
}

func ParseData(lines []string) []message {
	data := make([]message, 0)
	re := regexp.MustCompile(`^(.*) \| (.*)$`)
	for _, line := range lines {

		match1 := re.FindStringSubmatch(line)
		input := strings.Fields(match1[1])
		output := strings.Fields(match1[2])
		message := message{input: input, output: output}
		data = append(data, message)
	}
	return data
}

func Step1_CountOccurrenceOfSpecialDigitsInOutput(messages []message) int {
	count := 0
	for _, message := range messages {
		for _, digit := range message.output {
			if len(digit) == 2 {
				// 1
				count += 1
			} else if len(digit) == 3 {
				// 7
				count += 1
			} else if len(digit) == 4 {
				// 4
				count += 1
			} else if len(digit) == 7 {
				// 8
				count += 1
			}
		}
	}
	return count
}

// ##############################################################

type remainningPossibilities map[rune]string

// func getNewGuess() remainningPossibilities {
// 	return remainningPossibilities{
// 		'a': "ABCDEFG",
// 		'b': "ABCDEFG",
// 		'c': "ABCDEFG",
// 		'd': "ABCDEFG",
// 		'e': "ABCDEFG",
// 		'f': "ABCDEFG",
// 		'g': "ABCDEFG",
// 	}
// }

func Step2_GessDigits(messages []message) int {
	count := 0
	for _, message := range messages {
		guess := guessByOccurrence(message)
		displayNumber := ""
		for _, digit := range message.output {
			number := readDigit(digit, guess)
			displayNumber = fmt.Sprint(displayNumber, number)
		}

		displayedValue, _ := strconv.Atoi(displayNumber)
		count += displayedValue
	}
	return count
}

//  AAAA
// B    C
// B    C
//  DDDD
// E    F
// E    F
//  GGGG

func guessByOccurrence(message message) map[rune]string {

	guess := make(map[rune]string)

	patternFor1 := ""
	for _, pattern := range message.input {
		if len(pattern) == 2 {
			patternFor1 = pattern
			break
		}
	}

	patternFor4 := ""
	for _, pattern := range message.input {
		if len(pattern) == 4 {
			patternFor4 = pattern
			break
		}
	}

	occurrences := map[rune]int{}
	for _, pattern := range message.input {
		for _, char := range pattern {
			occurrences[char] += 1
		}
	}

	// print occurences
	// for k, v := range occurrences {
	// 	fmt.Println(string(k), ": ", v)
	// }

	for char, count := range occurrences {
		// segment E only occurs 4 times (2,6,8,0)
		if count == 4 {
			guess[char] = "E"
		}
		// segment B occurs 6 times
		if count == 6 {
			guess[char] = "B"
		}
		// segment F occurs 9 times
		if count == 9 {
			guess[char] = "F"
		}
		// segment C occurs 8 times and is used by 1
		// segment A occurs 8 times but is not used by 1
		if count == 8 && strings.Contains(patternFor1, string(char)) {
			guess[char] = "C"
		} else if count == 8 {
			guess[char] = "A"
		}
		// segment D occurs 7 times and is used by 4
		// segment G occurs 7 times but is not used by 4
		if count == 7 && strings.Contains(patternFor4, string(char)) {
			guess[char] = "D"
		} else if count == 7 {
			guess[char] = "G"
		}
	}
	return guess
}

func displayRemaining(guess remainningPossibilities) {
	for key, value := range guess {
		fmt.Println(string(key), " : ", value)
	}
}

func readDigit(pattern string, guess remainningPossibilities) string {
	Digits := map[string]string{
		"CF":      "1",
		"ACDEG":   "2",
		"ACDFG":   "3",
		"BCDF":    "4",
		"ABDFG":   "5",
		"ABDEFG":  "6",
		"ACF":     "7",
		"ABCDEFG": "8",
		"ABCDFG":  "9",
		"ABCEFG":  "0",
	}

	translatedPattern := ""
	for _, char := range pattern {
		translatedChar := guess[char]
		translatedPattern = fmt.Sprint(translatedPattern, translatedChar)
	}
	return Digits[SortString(translatedPattern)]
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
