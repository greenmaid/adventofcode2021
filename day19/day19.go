package day19

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type scanner struct {
	name    string
	beacons [][3]int
}

func ParseData(input []string) []scanner {
	re1 := regexp.MustCompile(`^--- scanner (\d+) ---$`)
	re2 := regexp.MustCompile(`^(-?\d+),(-?\d+),(-?\d+)$`)

	currentScanner := scanner{name: "xxx"}
	scanners := make([]scanner, 0)

	for _, line := range input {
		match1 := re1.FindStringSubmatch(line)
		if len(match1) == 2 {
			if currentScanner.name != "xxx" {
				scanners = append(scanners, currentScanner)
			}
			currentScanner = scanner{name: match1[1]}
		} else {
			match2 := re2.FindStringSubmatch(line)
			if len(match2) == 4 {
				x, _ := strconv.Atoi(match2[1])
				y, _ := strconv.Atoi(match2[2])
				z, _ := strconv.Atoi(match2[3])
				currentScanner.beacons = append(currentScanner.beacons, [3]int{x, y, z})
			}
		}
	}
	scanners = append(scanners, currentScanner)
	return scanners
}

func Run(scanners []scanner) (int, int) {

	scannersCoords := make([][3]int, 0)
	scannersCoords = append(scannersCoords, [3]int{0, 0, 0})

	for {
		if len(scanners) == 1 {
			break
		}
		for i := 1; i < len(scanners); i++ {
			matches := findMatches(scanners[0], scanners[i])
			if len(matches) >= 12 {
				coord, scanInOtherRef := getScannerTranslation(scanners[0], scanners[i], matches)
				scanners[0] = mergeScanners(scanners[0], scanInOtherRef)
				scanners = removeScanner(scanners, i)
				scannersCoords = append(scannersCoords, coord)
			}
		}
	}
	return len(scanners[0].beacons), maxDistance(scannersCoords)
}

func maxDistance(scannersCoords [][3]int) int {
	maxDistance := 0
	for _, scan1 := range scannersCoords {
		for _, scan2 := range scannersCoords {
			distance := abs(scan2[0]-scan1[0]) + abs(scan2[1]-scan1[1]) + abs(scan2[2]-scan1[2])
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}
	return maxDistance
}

func getBeaconDistance(beacon1 [3]int, beacon2 [3]int) float64 {
	x := beacon2[0] - beacon1[0]
	y := beacon2[1] - beacon1[1]
	z := beacon2[2] - beacon1[2]
	return math.Sqrt(float64(x*x + y*y + z*z))
}

func getDistanceList(beacon [3]int, scanner scanner) []float64 {
	distances := make([]float64, 0)
	for _, b := range scanner.beacons {
		distances = append(distances, getBeaconDistance(beacon, b))
	}
	return distances
}

func findPossiblematches(beacon [3]int, currentScanner scanner, scanner scanner) ([]int, int) {
	scannnerDistances := make([][]float64, 0)
	for _, b := range scanner.beacons {
		scannnerDistances = append(scannnerDistances, getDistanceList(b, scanner))
	}
	maxMatching := 0
	possibleMatch := make([]int, 0)
	currentDistances := getDistanceList(beacon, currentScanner)
	for index, scannerDistance := range scannnerDistances {
		matchings := countMatchingElements(currentDistances, scannerDistance)
		if matchings > maxMatching {
			maxMatching = matchings
			possibleMatch = []int{index}
		} else if matchings == maxMatching {
			possibleMatch = append(possibleMatch, index)
		}
	}

	return possibleMatch, maxMatching
}

func countMatchingElements(list1 []float64, list2 []float64) int {
	count := 0
	for _, val1 := range list1 {
		for index2, val2 := range list2 {
			if val1 == val2 {
				count += 1
				list2 = remove(list2, index2)
				break
			}
		}
	}

	return count
}

func remove(s []float64, i int) []float64 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeScanner(s []scanner, i int) []scanner {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func findMatches(scanner1 scanner, scanner2 scanner) [][2]int {
	highProbabilityMatchCount := 0
	highProbabilityMatch := make([][2]int, 0)
	highProbabilityUniqueMatchCount := 0
	for index, beacon := range scanner1.beacons {
		possible, matchCount := findPossiblematches(beacon, scanner1, scanner2)
		if matchCount >= 12 {
			highProbabilityMatchCount += 1
			if len(possible) == 1 {
				highProbabilityUniqueMatchCount += 1
				// fmt.Println("Scanner", scanner1.name, "beacon", beacon, "matches with scanner", scanner2.name, "beacon", scanner2.beacons[possible[0]], "with", matchCount, "overlapping")
				highProbabilityMatch = append(highProbabilityMatch, [2]int{index, possible[0]})
			}
		}
	}

	return highProbabilityMatch
}

func getScannerTranslation(scannerA scanner, scannerB scanner, matches [][2]int) ([3]int, scanner) {

	rotationMatrixAtoB := make(map[int]string)
	for x := 0; x <= len(matches)-2; x++ {
		point0A := scannerA.beacons[matches[x][0]]
		point1A := scannerA.beacons[matches[x+1][0]]
		point0B := scannerB.beacons[matches[x][1]]
		point1B := scannerB.beacons[matches[x+1][1]]

		diffViewedByA := [3]int{point1A[0] - point0A[0], point1A[1] - point0A[1], point1A[2] - point0A[2]}
		diffViewedByB := [3]int{point1B[0] - point0B[0], point1B[1] - point0B[1], point1B[2] - point0B[2]}

		if abs(diffViewedByA[0]) == abs(diffViewedByA[1]) ||
			abs(diffViewedByA[1]) == abs(diffViewedByA[2]) ||
			abs(diffViewedByA[2]) == abs(diffViewedByA[0]) {
			// couple of points is not good, try another
			continue
		}
		matrix := [3]string{"x", "y", "z"}
		matrixOpp := [3]string{"-x", "-y", "-z"}
		for indexA, valA := range diffViewedByA {
			for indexB, valB := range diffViewedByB {
				if abs(valA) == abs(valB) {
					if valA == valB {
						rotationMatrixAtoB[indexA] = matrix[indexB]
					} else {
						rotationMatrixAtoB[indexA] = matrixOpp[indexB]
					}
				}
			}
		}
		break
	}

	point0A := scannerA.beacons[matches[0][0]]
	point0B := scannerB.beacons[matches[0][1]]
	point0BOriented := [3]int{0}
	for i := 0; i < 3; i++ {
		switch rotationMatrixAtoB[i] {
		case "x":
			point0BOriented[i] = point0B[0]
		case "-x":
			point0BOriented[i] = (-1) * point0B[0]
		case "y":
			point0BOriented[i] = point0B[1]
		case "-y":
			point0BOriented[i] = (-1) * point0B[1]
		case "z":
			point0BOriented[i] = point0B[2]
		case "-z":
			point0BOriented[i] = (-1) * point0B[2]
		}
	}

	coordScannerBInARef := [3]int{0}
	coordScannerBInARef[0] = point0A[0] - point0BOriented[0]
	coordScannerBInARef[1] = point0A[1] - point0BOriented[1]
	coordScannerBInARef[2] = point0A[2] - point0BOriented[2]

	scannerBInARef := scanner{name: fmt.Sprint(scannerB.name, "with", scannerA.name, "referentiel")}
	for _, value := range scannerB.beacons {
		b := [3]int{0}
		for i := 0; i < 3; i++ {
			switch rotationMatrixAtoB[i] {
			case "x":
				b[i] = value[0] + coordScannerBInARef[i]
			case "-x":
				b[i] = ((-1) * value[0]) + coordScannerBInARef[i]
			case "y":
				b[i] = value[1] + coordScannerBInARef[i]
			case "-y":
				b[i] = ((-1) * value[1]) + coordScannerBInARef[i]
			case "z":
				b[i] = value[2] + coordScannerBInARef[i]
			case "-z":
				b[i] = ((-1) * value[2]) + coordScannerBInARef[i]
			}
		}
		scannerBInARef.beacons = append(scannerBInARef.beacons, b)
	}

	return coordScannerBInARef, scannerBInARef
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func mergeScanners(scanner1 scanner, scanner2 scanner) scanner {
	resultScanner := scanner{name: scanner1.name}
	resultScanner.beacons = scanner1.beacons
	for _, beacon := range scanner2.beacons {
		if !isInList(beacon, scanner1.beacons) {
			resultScanner.beacons = append(resultScanner.beacons, beacon)
		}
	}
	return resultScanner
}

func isInList(beacon [3]int, beaconList [][3]int) bool {
	for _, b := range beaconList {
		if b == beacon {
			return true
		}
	}
	return false
}
