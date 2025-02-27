package utilities

func FindPaths(graph Graph) [][]string {
	var paths [][]string
	visited := make(map[string]struct{})
	var currentPath []string

	// Start DFS from the start room
	findPathsDFS(graph.Start, graph.End, visited, &currentPath, &paths)

	return paths
}

func findPathsDFS(current, end *Room, visited map[string]struct{}, currentPath *[]string, paths *[][]string) {
	// Mark current room as visited
	visited[current.Name] = struct{}{}
	*currentPath = append(*currentPath, current.Name)

	// If we reached the end room, add the current path to our results
	if current == end {
		*paths = append(*paths, append([]string(nil), *currentPath...))
	} else {
		// Explore all neighbors
		for _, neighbor := range current.Neighbors {
			if _, seen := visited[neighbor.Name]; !seen {
				findPathsDFS(neighbor, end, visited, currentPath, paths)
			}
		}
	}

	// Backtrack: remove current room from path and mark as unvisited
	delete(visited, current.Name)
	*currentPath = (*currentPath)[:len(*currentPath)-1]
}
