package utilities

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Represents a single room in the colony
type Room struct {
	Name      string
	Neighbors []*Room
	IsStart   bool
	IsEnd     bool
}

// Represents an ant moving through the colony
type Ant struct {
	ID   int
	Path []string // Room in the ant path
}

// Represents the ant farm
type Graph struct {
	Rooms map[string]*Room
	Start *Room
	End   *Room
}

func ParseInput(file string) (int, *Graph, error) {
	graph := &Graph{Rooms: make(map[string]*Room)} // Create a graph to store rooms and connections (farm)
	var ants int                                   // Store the number of ants

	f, err := os.Open(file)
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	phase := "ants" // First phase is parsing the number of ants
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing whitespace
		if strings.HasPrefix(line, "#") {         // Ignore comments
			if line == "##start" || line == "##end" { // Rroom phases (##start or ##end)

				phase = line
			}
			continue
		}

		switch phase {
		case "ants":
			fmt.Sscanf(line, "%d", &ants) // Parse the number of ants
			phase = "rooms"               // Go to parsing rooms
		case "rooms":
			if strings.Contains(line, "-") { // If the line contains "-", it's a link --> switch to link parsing
				phase = "links"
				parts := strings.Split(line, "-")
				if len(parts) != 2 {
					return 0, nil, errors.New("ERROR: Invalid link format")
				}
				room1, room2 := parts[0], parts[1]
				//add come and go conections
				graph.Rooms[room1].Neighbors = append(graph.Rooms[room1].Neighbors, graph.Rooms[room2])
				graph.Rooms[room2].Neighbors = append(graph.Rooms[room2].Neighbors, graph.Rooms[room1])
			} else { // Parse a room
				parts := strings.Fields(line) // Split the line into parts
				if len(parts) != 3 {
					return 0, nil, errors.New("ERROR: invalid room format")
				}
				name := parts[0]
				graph.Rooms[name] = &Room{Name: name} // Create a room and add it to the graph

				if phase == "##start" {
					graph.Start = graph.Rooms[name]
					graph.Start.IsStart = true
					phase = "rooms0"
				} else if phase == "##end" {
					graph.End = graph.Rooms[name]
					graph.End.IsEnd = true
					phase = "rooms"
				}
			}

		case "links": // Parse link

			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return 0, nil, errors.New("ERROR: invalid link format")
			}
			room1, room2 := parts[0], parts[1]
			//add  conections
			graph.Rooms[room1].Neighbors = append(graph.Rooms[room1].Neighbors, graph.Rooms[room2])
			graph.Rooms[room2].Neighbors = append(graph.Rooms[room2].Neighbors, graph.Rooms[room1])
		}

	}

	if graph.Start == nil || graph.End == nil { // chack start and end
		return 0, nil, errors.New("ERROR: missing start or end room")
	}

	return ants, graph, nil
}
