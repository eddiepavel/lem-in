package utilities

import (
	"sort"
)

func FilterPaths(paths [][]string, graph *Graph) [][]string {
	// Sort the paths by length
	sortPaths(paths)

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
