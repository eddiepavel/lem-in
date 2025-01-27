package utilities

func FilterPaths(paths [][]string, graph *Graph) [][]string {
	// A slice to store the filtered paths
	newPaths := make([][]string, 0)

	// A map to track visited rooms
	visitedRooms := make(map[string]bool)

	// Iterate over each path
	for _, path := range paths {
		// A flag to indicate if the path should be added
		appendPath := true

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

			// Mark the room as visited
			visitedRooms[room] = true
		}

		// Add the path to the result if it is valid
		if appendPath {
			newPaths = append(newPaths, path)
		}
	}

	return newPaths
}
