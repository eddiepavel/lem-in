package utilities

import (
	"sort"
)

func FilterPaths(paths [][]string, graph *Graph, antCount *int) [][]string {
	if len(paths) == 1 {
		return paths
	}
	// Sort the paths by length
	sortPaths(paths)

	// Filter paths to keep non-overlapping paths
	result := filterNonOverlappingPaths(paths, graph, nil, nil)
	result = filterNonOverlappingPaths(paths, graph, &result, antCount)

	return result
}

func filterNonOverlappingPaths(paths [][]string, graph *Graph, result *[][]string, ants *int) [][]string {
	if result != nil {
		paths = removeShortestPaths(paths)
	}

	newPaths := make([][]string, 0)
	visitedRooms := make(map[string]struct{})

	for _, path := range paths {
		if isValidPath(path, graph, visitedRooms) {
			markRoomsAsVisited(path, graph, visitedRooms)
			newPaths = append(newPaths, path)
		}
	}
	if result != nil && len(newPaths) > 0 {
		return compareAntThroughput(newPaths, *result, *ants, graph)
	}

	return newPaths
}

func removeShortestPaths(paths [][]string) [][]string {
	shortestLength := len(paths[0])
	filteredPaths := make([][]string, 0)
	for _, path := range paths {
		if len(path) > shortestLength {
			filteredPaths = append(filteredPaths, path)
		}
	}
	if len(filteredPaths) == 0 {
		return paths
	}
	return filteredPaths
}

func isValidPath(path []string, graph *Graph, visitedRooms map[string]struct{}) bool {
	for _, room := range path {
		if room == graph.Start.Name || room == graph.End.Name {
			continue
		}
		if _, visited := visitedRooms[room]; visited {
			return false
		}
	}
	return true
}

func markRoomsAsVisited(path []string, graph *Graph, visitedRooms map[string]struct{}) {
	for _, room := range path {
		if room != graph.Start.Name && room != graph.End.Name {
			visitedRooms[room] = struct{}{}
		}
	}
}

func sortPaths(paths [][]string) {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

func compareAntThroughput(paths1, paths2 [][]string, antsCount int, graph *Graph) [][]string {
	_, steps1 := MoveAnts(paths1, antsCount, graph)
	_, steps2 := MoveAnts(paths2, antsCount, graph)
	if steps1 < steps2 {
		return paths1
	}
	return paths2
}
