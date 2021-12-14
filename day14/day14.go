package day14

import (
	"fmt"
	"regexp"
	"sort"
)

type Polymer struct {
	start     byte
	end       byte
	pairs     map[[2]byte]int
	elemCount map[byte]int
}

func ParseData(data []string) (Polymer, map[[2]byte]byte) {

	polymer := data[0]
	pairs := make(map[[2]byte]int)
	count := make(map[byte]int)
	for i := 0; i < len(polymer); i++ {
		count[polymer[i]] += 1
		if i < len(polymer)-1 {
			pair := [2]byte{polymer[i], polymer[i+1]}
			pairs[pair] += 1
		}
	}

	initPolymer := Polymer{
		start:     polymer[0],
		end:       polymer[len(polymer)-1],
		pairs:     pairs,
		elemCount: count,
	}

	polymerizationRules := make(map[[2]byte]byte)

	re := regexp.MustCompile(`^(\w)(\w) -> (\w)$`)
	for i := 2; i < len(data); i++ {
		match := re.FindStringSubmatch(data[i])
		polymerizationRules[[2]byte{match[1][0], match[2][0]}] = match[3][0]
	}
	return initPolymer, polymerizationRules
}

func Step1(polymer Polymer, rules map[[2]byte]byte) int {
	currentPolymer := polymer
	for step := 1; step <= 10; step++ {
		nextPolymer := polymerization(currentPolymer, rules)
		currentPolymer = nextPolymer
	}
	max, min := computeMostCommonAndLessCommonElements(currentPolymer)
	return max - min
}

func Step2(polymer Polymer, rules map[[2]byte]byte) int {
	currentPolymer := polymer
	for step := 1; step <= 40; step++ {
		nextPolymer := polymerization(currentPolymer, rules)
		currentPolymer = nextPolymer
	}
	max, min := computeMostCommonAndLessCommonElements(currentPolymer)
	return max - min
}

func polymerization(polymer Polymer, rules map[[2]byte]byte) Polymer {

	newPairs := make(map[[2]byte]int)
	for pair, count := range polymer.pairs {
		if addElem, ok := rules[pair]; ok {
			newPairs[[2]byte{pair[0], addElem}] += count
			newPairs[[2]byte{addElem, pair[1]}] += count
			polymer.elemCount[addElem] += count
		} else {
			newPairs[pair] += count
		}
	}
	polymer.pairs = newPairs
	return polymer
}

func computeMostCommonAndLessCommonElements(currentPolymer Polymer) (int, int) {

	elementCountList := make([]int, 0)
	for _, value := range currentPolymer.elemCount {
		elementCountList = append(elementCountList, value)
	}
	sort.Ints(elementCountList)
	// fmt.Println(elementCountList)

	return elementCountList[len(elementCountList)-1], elementCountList[0]
}

func alternativeCount(polymer Polymer) (int, int) {
	elementRepartition := make(map[byte]int)
	elementRepartition[polymer.start] = 1
	elementRepartition[polymer.end] = 1
	for pair, count := range polymer.pairs {
		elementRepartition[pair[0]] += count / 2
		elementRepartition[pair[1]] += count / 2
	}

	elementCountList := make([]int, 0)
	for _, value := range polymer.elemCount {
		elementCountList = append(elementCountList, value)
	}
	sort.Ints(elementCountList)
	fmt.Println(elementCountList)

	return elementCountList[len(elementCountList)-1], elementCountList[0]
}
