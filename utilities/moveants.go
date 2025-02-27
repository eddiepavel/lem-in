package utilities

import (
	"fmt"
	"strings"
)

func MoveAnts(paths [][]string, antsCount int, graph *Graph) (string, int) {
	ants := initializeAnts(paths, antsCount)
	progress := make([]byte, antsCount)
	finished := 0
	roomOccupancy := make(map[string]bool)
	var outputBuilder strings.Builder
	steps := 0

	for turn := 1; finished < antsCount; turn++ {
		output := ""
		pathsUsed := make([]bool, len(paths))

		for i := 0; i < antsCount; i++ {
			if progress[i] >= byte(len(ants[i].Path)-1) {
				continue
			}

			currentRoom, nextRoom := ants[i].Path[progress[i]], ants[i].Path[progress[i]+1]

			if currentRoom == graph.Start.Name && !canUsePath(i, len(paths), pathsUsed, antsCount, paths) {
				continue
			}

			if isRoomAvailable(nextRoom, roomOccupancy, graph) {
				updateRoomOccupancy(currentRoom, nextRoom, roomOccupancy, graph, ants[i].ID)
				progress[i]++
				output += fmt.Sprintf("L%d-%s ", ants[i].ID, nextRoom)

				if progress[i] == byte(len(ants[i].Path)-1) {
					finished++
				}
			}
		}

		if output != "" {
			outputBuilder.WriteString(strings.TrimSpace(output))
			outputBuilder.WriteString("\n")
			steps++
		}
	}

	return strings.TrimSuffix(outputBuilder.String(), "\n"), steps
}

func initializeAnts(paths [][]string, antsCount int) []Ant {
	ants := make([]Ant, antsCount)
	for i := 0; i < antsCount; i++ {
		pathIndex := i % len(paths)
		if len(paths) != 1 && i == antsCount-1 && len(paths[0]) != len(paths[pathIndex]) {
			pathIndex = 0
		}
		ants[i] = Ant{ID: i + 1, Path: paths[pathIndex]}
	}
	return ants
}

func canUsePath(i, pathsLen int, pathsUsed []bool, antsCount int, paths [][]string) bool {
	pathIndex := i % pathsLen
	if i == antsCount-1 && pathsLen != 1 && len(paths[0]) != len(paths[pathIndex]) {
		pathIndex = 0
	}
	if pathsUsed[pathIndex] {
		return false
	}
	pathsUsed[pathIndex] = true
	return true
}

func isRoomAvailable(nextRoom string, roomOccupancy map[string]bool, graph *Graph) bool {
	return !roomOccupancy[nextRoom] || nextRoom == graph.End.Name
}

func updateRoomOccupancy(currentRoom, nextRoom string, roomOccupancy map[string]bool, graph *Graph, antID int) {
	if currentRoom != "" && currentRoom != graph.End.Name {
		roomOccupancy[currentRoom] = false
	}
	if nextRoom != graph.End.Name {
		roomOccupancy[nextRoom] = true
	}
}
