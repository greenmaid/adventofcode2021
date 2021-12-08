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

func getNewGuess() remainningPossibilities {
	return remainningPossibilities{
		'a': "ABCDEFG",
		'b': "ABCDEFG",
		'c': "ABCDEFG",
		'd': "ABCDEFG",
		'e': "ABCDEFG",
		'f': "ABCDEFG",
		'g': "ABCDEFG",
	}
}
func Step2_GessDigits(messages []message) int {

	count := 0
	for _, message := range messages {
		guess := guess(message)
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

func guess(message message) remainningPossibilities {
	guess := getNewGuess()
	guess = findA(message, guess)
	guess = findG(message, guess)
	guess = guessByOccurrence(message, guess)
	guess = removeGuessedFromOtherPossibilities(guess)
	return guess
}

// A is only missing in 1 and 4
// A is the character which is contained by 7 and not 1
func findA(message message, guess remainningPossibilities) remainningPossibilities {
	possibilities := ""
	for _, pattern := range message.input {
		if len(pattern) == 3 {
			possibilities = pattern
			break
		}
	}
	for _, pattern := range message.input {
		if len(pattern) == 2 {
			for _, char := range pattern {
				possibilities = strings.ReplaceAll(possibilities, string(char), "")
			}
			break
		}
	}

	if len(possibilities) == 1 {
		for _, value := range possibilities {
			guess[value] = "A"
		}

	}
	return guess
}

// G is only missing in 1, 7 and 4
func findG(message message, guess remainningPossibilities) remainningPossibilities {
	possibilities := "abcdefg"
	for _, pattern := range message.input {
		if len(pattern) <= 4 {
			for _, char := range pattern {
				possibilities = strings.ReplaceAll(possibilities, string(char), "")
			}
			continue
		}
		if len(pattern) > 4 {
			for _, char := range "abcdefg" {
				if !strings.Contains(pattern, string(char)) {
					possibilities = strings.ReplaceAll(possibilities, string(char), "")
				}
			}
		}
	}

	if len(possibilities) == 1 {
		for _, value := range possibilities {
			guess[value] = "G"
		}

	}
	return guess
}

// When a segment is guessed, its possibility should be removed from others
func removeGuessedFromOtherPossibilities(guess remainningPossibilities) remainningPossibilities {
	for key := range guess {
		if len(guess[key]) == 1 {
			guessedSegment := guess[key]
			for k := range guess {
				if key != k {
					guess[k] = strings.ReplaceAll(guess[k], guessedSegment, "")
				}
			}
		}
	}
	return guess
}

func guessByOccurrence(message message, guess remainningPossibilities) remainningPossibilities {
	patternFor1 := ""
	for _, pattern := range message.input {
		if len(pattern) == 2 {
			patternFor1 = pattern
			break
		}
	}
	occurrences := map[rune]int{}
	for _, pattern := range message.input {
		for _, char := range pattern {
			occurrences[char] += 1
		}
	}

	for char, value := range occurrences {
		// segment E only occurs 4 times (2,6,8,0)
		if value == 4 {
			guess[char] = "E"
		}
		// segment B occurs 6 times
		if value == 6 {
			guess[char] = "B"
		}
		// segment F occurs 9 times
		if value == 9 {
			guess[char] = "F"
		}
		// C and F are used by 1
		// segment C occurs 8 times
		if value == 8 && strings.Contains(patternFor1, string(char)) {
			guess[char] = "C"
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
