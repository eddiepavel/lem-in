package utilities

import (
	"fmt"
	"strings"
)

func (g *Graph) MoveAnts() (string, int) {
	ants := g.initializeAnts()
	progress := make([]byte, g.Count)
	finished := 0
	roomOccupancy := make(map[string]bool)
	var outputBuilder strings.Builder
	steps := 0

	for turn := 1; finished < g.Count; turn++ {
		output := ""
		pathsUsed := make([]bool, len(g.FinalPaths))

		for i := 0; i < g.Count; i++ {
			if progress[i] >= byte(len(ants[i].Path)-1) {
				continue
			}

			currentRoom, nextRoom := ants[i].Path[progress[i]], ants[i].Path[progress[i]+1]

			if currentRoom == g.Start.Name && !g.canUsePath(i, pathsUsed) {
				continue
			}

			if g.isRoomAvailable(nextRoom, roomOccupancy) {
				g.updateRoomOccupancy(currentRoom, nextRoom, roomOccupancy)
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
	if g.FinalOutput == "" {
		g.FinalOutput = strings.TrimSuffix(outputBuilder.String(), "\n")
		g.FinalSteps = steps
		return "", 0
	}
	return strings.TrimSuffix(outputBuilder.String(), "\n"), steps
}

func (g *Graph) initializeAnts() []Ant {
	ants := make([]Ant, g.Count)
	for i := 0; i < g.Count; i++ {
		pathIndex := i % len(g.FinalPaths)
		if len(g.FinalPaths) != 1 && i == g.Count-1 && len(g.AllPaths[0]) != len(g.AllPaths[pathIndex]) {
			pathIndex = 0
		}
		ants[i] = Ant{ID: i + 1, Path: g.FinalPaths[pathIndex]}
	}
	return ants
}

func (g *Graph) canUsePath(i int, pathsUsed []bool) bool {
	pathIndex := i % len(g.FinalPaths)
	if i == g.Count-1 && len(g.AllPaths) != 1 && len(g.AllPaths[0]) != len(g.AllPaths[pathIndex]) {
		pathIndex = 0
	}
	if pathsUsed[pathIndex] {
		return false
	}
	pathsUsed[pathIndex] = true
	return true
}

func (g *Graph) isRoomAvailable(nextRoom string, roomOccupancy map[string]bool) bool {
	return !roomOccupancy[nextRoom] || nextRoom == g.End.Name
}

func (g *Graph) updateRoomOccupancy(currentRoom, nextRoom string, roomOccupancy map[string]bool) {
	if currentRoom != "" && currentRoom != g.End.Name {
		roomOccupancy[currentRoom] = false
	}
	if nextRoom != g.End.Name {
		roomOccupancy[nextRoom] = true
	}
}
