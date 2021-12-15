package day15

import (
	"adventofcode2021/common"

	"github.com/yourbasic/graph"
)

func ParseData(input []string, factor int) (*graph.Mutable, int) {
	grid := getGrid(input, factor)

	// Create graph from grid
	// simple graph needs int as node referece
	//     => using x+1000y  (orthogonality working as long x<1000)
	g := graph.New(len(grid) * 1000)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			current := x + 1000*y
			if x > 0 {
				g.AddCost(current, current-1, int64(grid[y][x-1]))
			}
			if x < len(grid[0])-1 {
				g.AddCost(current, current+1, int64(grid[y][x+1]))
			}
			if y > 0 {
				g.AddCost(current, current-1000, int64(grid[y-1][x]))
			}
			if y < len(grid)-1 {
				g.AddCost(current, current+1000, int64(grid[y+1][x]))
			}
		}

	}
	endNodeIndex := len(grid[0]) - 1 + 1000*(len(grid)-1)
	return g, endNodeIndex
}

func getGrid(input []string, factor int) [][]int {
	grid := make([][]int, 0)
	for _, line := range input {
		grid = append(grid, common.ParseLineAsBits(line))
	}

	expandedGrid := make([][]int, 0)
	for y := 0; y < len(grid)*factor; y++ {
		line := make([]int, 0)
		for x := 0; x < len(grid[0])*factor; x++ {
			value := grid[y%len(grid)][x%len(grid[0])] + y/len(grid) + x/len(grid[0])
			if value > 9 {
				value = value%10 + 1
			}
			line = append(line, value)
		}
		expandedGrid = append(expandedGrid, line)
	}

	// for _,line := range expandedGrid {
	// 	fmt.Println(line)
	// }
	return expandedGrid
}

func Run(g graph.Iterator, end int) int64 {
	_, dist := graph.ShortestPath(g, 0, end)
	// fmt.Println(path)
	return dist
}
