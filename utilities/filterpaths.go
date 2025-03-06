package utilities

import (
	"fmt"
	"sort"
)

func (g *Graph) FilterPaths() {
	// If there is only one path, no need to filter
	if len(g.AllPaths) == 1 {
		g.FinalPaths = g.AllPaths
		_, _ = g.MoveAnts()
		fmt.Println(g.FinalOutput)
		return
	}
	// Sort the paths by length
	g.sortPaths()

	// Filter paths to keep non-overlapping paths
	g.filterNonOverlappingPaths()
}

func (g *Graph) filterNonOverlappingPaths() {
	newPaths := make([][]string, 0)
	visitedRooms := make(map[string]struct{})

	for _, path := range g.AllPaths {
		if g.isValidPath(path, visitedRooms) {
			g.markRoomsAsVisited(path, visitedRooms)
			newPaths = append(newPaths, path)
		}
	}
	// If new combination of paths is not empty, compare throughput with previous combination
	if g.FinalPaths != nil && len(newPaths) > 0 {
		g.compareAntThroughput(newPaths)
		return
	}

	g.FinalPaths = newPaths
	_, _ = g.MoveAnts() // Move ants to get the first output
	g.BruteForce()      // Brute force to see if new combination of paths can be found
}

func (g *Graph) BruteForce() {
	copy := g.AllPaths
	g.removeShortestPaths()
	if len(g.AllPaths) == len(copy) { // If paths are not removed, means all paths have same length
		fmt.Println(g.FinalOutput)
		return
	}
	g.filterNonOverlappingPaths() // Filter paths again
}

func (g *Graph) removeShortestPaths() {
	shortestLength := len(g.AllPaths[0])
	filteredPaths := make([][]string, 0)
	for _, path := range g.AllPaths {
		if len(path) > shortestLength {
			filteredPaths = append(filteredPaths, path)
		}
	}
	if len(filteredPaths) == 0 {
		return
	}
	g.AllPaths = filteredPaths
}

func (g *Graph) isValidPath(path []string, visitedRooms map[string]struct{}) bool {
	for _, room := range path {
		if room == g.Start.Name || room == g.End.Name {
			continue
		}
		if _, visited := visitedRooms[room]; visited {
			return false
		}
	}
	return true
}

func (g *Graph) markRoomsAsVisited(path []string, visitedRooms map[string]struct{}) {
	for _, room := range path {
		if room != g.Start.Name && room != g.End.Name {
			visitedRooms[room] = struct{}{}
		}
	}
}

func (g *Graph) sortPaths() {
	sort.Slice(g.AllPaths, func(i, j int) bool {
		return len(g.AllPaths[i]) < len(g.AllPaths[j])
	})
}

func (g *Graph) compareAntThroughput(newPaths [][]string) {
	output1, steps1 := g.FinalOutput, g.FinalSteps
	g.FinalPaths = newPaths
	output2, steps2 := g.MoveAnts()
	if steps1 < steps2 {
		fmt.Println(output1)
		return
	}
	fmt.Println(output2)
}
