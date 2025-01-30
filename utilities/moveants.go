package utilities

import (
	"fmt"
	"strings"
)

func MoveAnts(paths [][]string, antsCount int, graph *Graph) (string, int) {
	// Initialize ants with their paths
	ants := make([]Ant, antsCount)
	for i := 0; i < antsCount; i++ {
		if len(paths) != 1 {
			if i == antsCount-1 {
				ants[i] = Ant{
					ID:   i + 1,
					Path: paths[0],
				}
				continue
			}
		}
		ants[i] = Ant{
			ID:   i + 1,
			Path: paths[i%len(paths)],
		}
	}

	progress := make([]int, antsCount)    // Track current step in each ant's path
	finished := 0                         // Count of ants that have finished
	roomOccupancy := make(map[string]int) // Track room occupancy, excluding the end room

	var outputBuilder strings.Builder
	steps := 0 // Step counter

	for turn := 1; finished < antsCount; turn++ {
		output := ""
		pathsUsed := make(map[int]bool) // Tracks paths that have used their start room allowance this turn

		for i := 0; i < antsCount; i++ {
			if progress[i] >= len(ants[i].Path)-1 {
				continue // Ant has already finished
			}

			currentRoom := ants[i].Path[progress[i]]
			nextRoom := ants[i].Path[progress[i]+1]

			// Check if the ant is in the start room and handle path allowance
			if currentRoom == graph.Start.Name {
				pathIndex := i % len(paths)
				if i == antsCount-1 && len(paths) != 1 {
					pathIndex = 0
				}
				if pathsUsed[pathIndex] {
					continue // This path has already sent an ant this turn
				}
				pathsUsed[pathIndex] = true
			}

			// Check if the next room is available (end room is always allowed)
			if roomOccupancy[nextRoom] == 0 || nextRoom == graph.End.Name {
				// Free the current room if it's not the end room
				if currentRoom != "" && currentRoom != graph.End.Name {
					roomOccupancy[currentRoom] = 0
				}

				// Occupy the next room if it's not the end room
				if nextRoom != graph.End.Name {
					// Check if next room is actually free now
					if roomOccupancy[nextRoom] != 0 {
						continue // Skip if another ant has already taken this room in the same turn
					}
					roomOccupancy[nextRoom] = ants[i].ID
				}

				progress[i]++
				output += fmt.Sprintf("L%d-%s ", ants[i].ID, nextRoom)

				// Check if the ant has finished
				if progress[i] == len(ants[i].Path)-1 {
					finished++
				}
			}
		}

		// Collect the turn's movements if any
		if output != "" {
			outputBuilder.WriteString(output[:len(output)-1]) // Trim trailing space
			outputBuilder.WriteString("\n")
			steps++ // Increment step counter
		}
	}
	// Trim the final newline if it exists
	finalOutput := outputBuilder.String()
	if len(finalOutput) > 0 && finalOutput[len(finalOutput)-1] == '\n' {
		finalOutput = finalOutput[:len(finalOutput)-1]
	}

	return finalOutput, steps
}
