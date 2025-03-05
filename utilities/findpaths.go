package utilities

func (g *Graph) FindPaths() {

	visited := make(map[string]struct{})
	var currentPath []string

	// Start DFS from the start room
	g.findPathsDFS(g.Start, visited, &currentPath)

}

func (g *Graph) findPathsDFS(current *Room, visited map[string]struct{}, currentPath *[]string) {
	// Mark current room as visited
	visited[current.Name] = struct{}{}
	*currentPath = append(*currentPath, current.Name)

	// If we reached the end room, add the current path to our results
	if current == g.End {
		g.Paths = append(g.Paths, append([]string(nil), *currentPath...))
	} else {
		// Explore all neighbors
		for _, neighbor := range current.Neighbors {
			if _, seen := visited[neighbor.Name]; !seen {
				g.findPathsDFS(neighbor, visited, currentPath)
			}
		}
	}

	// Backtrack: remove current room from path and mark as unvisited
	delete(visited, current.Name)
	*currentPath = (*currentPath)[:len(*currentPath)-1]
}
