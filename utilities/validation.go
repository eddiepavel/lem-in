package utilities

import (
	"errors"
	"fmt"
)

func ValidateGraph(graph *Graph) error {
	if graph.Start == nil {
		return errors.New("graph is missing a start room")
	}
	if graph.End == nil {
		return errors.New("graph is missing a end room")
	}

	queue := []*Room{graph.Start}    // Start from the Start room
	visited := make(map[string]bool) // Track visited rooms

	// BFS Loop
	for len(queue) > 0 {
		current := queue[0] // Dequeue the first room
		queue = queue[1:]   // Remove it from the queue

		// Mark the current room as visited
		visited[current.Name] = true

		// Check if we reached the End room
		if current == graph.End {
			return nil // Successfully found a path
		}

		// Enqueue all unvisited neighbors
		for _, neighbor := range current.Neighbors {
			if !visited[neighbor.Name] {
				queue = append(queue, neighbor)
				visited[neighbor.Name] = true // Mark as visited when enqueued
			}
		}
	}

	// If BFS completes without finding the End room
	return fmt.Errorf("no valid path from start to end")

}
