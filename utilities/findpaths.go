package utilities

func FindPaths(graph Graph) [][]string {
	paths := make([][]string, 0)
	visited := make(map[string]bool)
	currentPath := make([]string, 0)

	// Start DFS from the start room
	findPathsDFS(graph.Start, graph.End, visited, &currentPath, &paths)

	return paths
}

func findPathsDFS(current *Room, end *Room, visited map[string]bool, currentPath *[]string, paths *[][]string) {
	// Mark current room as visited
	visited[current.Name] = true
	*currentPath = append(*currentPath, current.Name)

	// If we reached the end room, add the current path to our results
	if current == end {
		// Create a copy of the path to store
		pathCopy := make([]string, len(*currentPath))
		copy(pathCopy, *currentPath)
		*paths = append(*paths, pathCopy)
	} else {
		// Explore all neighbors
		for _, neighbor := range current.Neighbors {
			if !visited[neighbor.Name] {
				findPathsDFS(neighbor, end, visited, currentPath, paths)
			}
		}
	}

	// Backtrack: remove current room from path and mark as unvisited
	visited[current.Name] = false
	*currentPath = (*currentPath)[:len(*currentPath)-1]
}
