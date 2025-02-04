package utilities

import (
	"sort"
)

var beenCalled bool

func FilterPaths(paths [][]string, graph *Graph, antcount int) [][]string {
	paths1 := FilterHelper(paths, graph) // keep non-overlapping paths starting from the shortest
	paths2 := FilterHelper(paths, graph) // keep non-overlapping paths skipping the shortest to brute force different combos
	result := paths1
	if len(paths1) != len(paths2) && len(paths2) > 0 {
		result = CompareAntThroughput(paths1, paths2, antcount, graph) // compare ant throughput for different sets of filtered paths
	}
	return result
}

func FilterHelper(paths [][]string, graph *Graph) [][]string {
	// Sort the paths by length
	sortPaths(paths)

	if beenCalled {
		// Remove paths that have length equal to the shortest one
		shortestLength := len(paths[0])
		filteredPaths := make([][]string, 0)
		for _, path := range paths {
			if len(path) > shortestLength {
				filteredPaths = append(filteredPaths, path)
			}
		}
		paths = filteredPaths
	} else {
		beenCalled = true
	}
	// A slice to store the filtered paths
	newPaths := make([][]string, 0)

	// A map to track visited rooms
	visitedRooms := make(map[string]bool)

	// Iterate over each path
	for _, path := range paths {
		appendPath := true
		currentPathRooms := make([]string, 0)

		// Check each room in the path
		for _, room := range path {
			// Skip start and end rooms
			if room == graph.Start.Name || room == graph.End.Name {
				continue
			}

			// If the room is already visited, mark as invalid
			if visitedRooms[room] {
				appendPath = false
				break
			}
			currentPathRooms = append(currentPathRooms, room)
		}

		// Add the path to the result if it is valid
		if appendPath {
			// Mark rooms as visited only if path is valid
			for _, room := range currentPathRooms {
				visitedRooms[room] = true
			}
			newPaths = append(newPaths, path)
		}
	}

	return newPaths
}

func sortPaths(paths [][]string) {
	// Sort the paths by length
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

// Compare ant throughput for different sets of filtered paths
func CompareAntThroughput(paths1, paths2 [][]string, antsCount int, graph *Graph) [][]string {
	_, steps1 := MoveAnts(paths1, antsCount, graph)
	_, steps2 := MoveAnts(paths2, antsCount, graph)

	if steps1 < steps2 {
		return paths1
	}
	return paths2
}
