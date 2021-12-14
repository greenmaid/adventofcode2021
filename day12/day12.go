package day12

import (
	"regexp"
	"strings"
)

type Cave struct {
	name  string
	big   bool
	links []string
}

func ParseDataToCaves(lines []string) map[string]Cave {
	linkMap := make(map[string][]string)
	re := regexp.MustCompile(`^(\w+)-(\w+)$`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		linkMap[match[1]] = append(linkMap[match[1]], match[2])
		linkMap[match[2]] = append(linkMap[match[2]], match[1])
	}
	caves := make(map[string]Cave)
	for name, links := range linkMap {
		big := false
		if name == strings.ToUpper(name) {
			big = true
		}
		caves[name] = Cave{name: name, big: big, links: links}

	}
	return caves
}

type Path []Cave

var paths []Path

func Step1_CountValidPaths(caves map[string]Cave) int {
	paths = make([]Path, 0)
	visitedTwiceSmallCave := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	path := Path{caves["start"]}
	visitCaves(caves, path, visitedTwiceSmallCave)
	return len(paths)
}

func Step2_CountValidPaths(caves map[string]Cave) int {
	paths = make([]Path, 0)
	visitedTwiceSmallCave := ""
	path := Path{caves["start"]}
	visitCaves(caves, path, visitedTwiceSmallCave)
	return len(paths)
}

func visitCaves(caves map[string]Cave, currentPath Path, visitedTwiceSmallCave string) {
	currentCave := currentPath[len(currentPath)-1]
	for _, next := range currentCave.links {
		nextCave := caves[next]
		if nextCave.name == "start" {
			continue
		}
		if nextCave.name == "end" {
			newPath := append(currentPath, nextCave)
			paths = append(paths, newPath)
			continue
		}
		if nextCave.big || !isInPath(nextCave, currentPath) {
			newPath := append(currentPath, nextCave)
			visitCaves(caves, newPath, visitedTwiceSmallCave)
			continue
		}
		if visitedTwiceSmallCave == "" {
			newPath := append(currentPath, nextCave)
			visitCaves(caves, newPath, nextCave.name)
			continue
		}
	}
}

func isInPath(cave Cave, path Path) bool {
	for _, caveInPath := range path {
		if caveInPath.name == cave.name {
			return true
		}
	}
	return false
}
