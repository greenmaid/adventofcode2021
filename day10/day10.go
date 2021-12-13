package day10

import "sort"

func Step1_CheckData(data []string) int {
	errorList := CheckData(data)
	return scoreErrors(errorList)
}

func Step2_CheckIncompleteData(data []string) int {
	errorList := CheckData(data)
	return scoreIncomplete(errorList)
}

func CheckData(data []string) []ChunkError {
	errorList := make([]ChunkError, 0)
	for _, line := range data {
		errorList = append(errorList, checkLine(line))
	}

	return errorList
}

type ChunkError struct {
	err            bool
	incomplete     bool
	errChar        rune
	finishingChars []rune
}

func checkLine(input string) ChunkError {

	chunkError := ChunkError{err: false, incomplete: false}
	openChunks := []rune("")

	for _, char := range input {
		if char == '(' || char == '[' || char == '<' || char == '{' {
			openChunks = append(openChunks, char)
		}
		if char == ')' || char == ']' || char == '>' || char == '}' {
			lastOpen := openChunks[len(openChunks)-1:][0]
			if char == ')' {
				if lastOpen == '(' {
					openChunks = openChunks[:len(openChunks)-1]
					continue
				}
			}
			if char == ']' {
				if lastOpen == '[' {
					openChunks = openChunks[:len(openChunks)-1]
					continue
				}
			}
			if char == '>' {
				if lastOpen == '<' {
					openChunks = openChunks[:len(openChunks)-1]
					continue
				}
			}
			if char == '}' {
				if lastOpen == '{' {
					openChunks = openChunks[:len(openChunks)-1]
					continue
				}
			}
			chunkError.err = true
			chunkError.errChar = char
			break
		}
	}

	if !chunkError.err {
		if len(openChunks) > 0 {
			chunkError.incomplete = true
			for _, char := range openChunks {
				if char == '(' {
					chunkError.finishingChars = append([]rune{')'}, chunkError.finishingChars...)
					continue
				}
				if char == '[' {
					chunkError.finishingChars = append([]rune{']'}, chunkError.finishingChars...)
					continue
				}
				if char == '<' {
					chunkError.finishingChars = append([]rune{'>'}, chunkError.finishingChars...)
					continue
				}
				if char == '{' {
					chunkError.finishingChars = append([]rune{'}'}, chunkError.finishingChars...)
					continue
				}
			}
		}
	}

	return chunkError
}

func scoreErrors(errors []ChunkError) int {
	score := 0
	for _, error := range errors {
		if error.err {
			if error.errChar == ')' {
				score += 3
			}
			if error.errChar == ']' {
				score += 57
			}
			if error.errChar == '}' {
				score += 1197
			}
			if error.errChar == '>' {
				score += 25137
			}
		}
	}
	return score
}

func scoreIncomplete(errors []ChunkError) int {
	scores := make([]int, 0)
	scoreMap := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	for _, error := range errors {
		score := 0
		if error.incomplete == true {
			for _, char := range error.finishingChars {
				score = score * 5
				score += scoreMap[char]
			}
			scores = append(scores, score)
		}
	}

	// get middle score
	sort.Ints(scores)
	middle := len(scores) / 2
	return scores[middle]

}
