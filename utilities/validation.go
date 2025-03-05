package utilities

import (
	"fmt"
)

func (g *Graph) ValidateGraph() error {
	if g == nil {
		return fmt.Errorf("graph is nil")
	}
	if g.Start == nil {
		return fmt.Errorf("graph start room is nil")
	}
	if g.End == nil {
		return fmt.Errorf("graph end room is nil")
	}

	queue := []*Room{g.Start}            // Start from the Start room
	visited := make(map[string]struct{}) // Track visited rooms

	// BFS Loop
	for len(queue) > 0 {
		current := queue[0] // Dequeue the first room
		queue = queue[1:]   // Remove it from the queue

		// Mark the current room as visited
		visited[current.Name] = struct{}{}

		// Check if we reached the End room
		if current == g.End {
			return nil // Successfully found a path
		}

		// Enqueue all unvisited neighbors
		for _, neighbor := range current.Neighbors {
			if _, seen := visited[neighbor.Name]; !seen {
				queue = append(queue, neighbor)
				visited[neighbor.Name] = struct{}{} // Mark as visited when enqueued
			}
		}
	}

	// If BFS completes without finding the End room
	return fmt.Errorf("no valid path from start to end")
}
