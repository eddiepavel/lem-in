package utilities

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Represents a single room in the colony
type Room struct {
	Name      string
	Neighbors []*Room
	X, Y      int
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
	var ants int
	var flag bool
	var StartingRooms int
	var EndingRooms int
	var ExtraRooms int
	var tunels int
	FirstFlag := true // Flag so we know the extra rooms befor ##StartRooms

	f, err := os.Open(file)
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()
	coordinatesMap := make(map[string]string)
	scanner := bufio.NewScanner(f)
	phase := "ants" // First phase is parsing the number of ants
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing whitespace
		//if strings.HasPrefix(line, "#") {          Ignore comments
		if line == "##start" { // Rroom phases (##start or ##end)
			FirstFlag = false // First time to start so we are done with extra rooms

			phase = "start-end"

		} else if line == "##end" {

			phase = "start-end"
		}

		//}

		switch phase {
		case "ants":

			_, err := fmt.Sscanf(line, "%d", &ants) // Parse the number of ants
			if err != nil || ants <= 0 {
				return 0, nil, errors.New("error: invalid number of ants")

			}

			phase = "rooms" // Go to parsing rooms
			continue

		case "rooms":

			if strings.Contains(line, "-") { // If the line contains "-", it's a link --> switch to link parsing
				tunels++

				parts := strings.Split(line, "-")
				if len(parts) != 2 {
					return 0, nil, errors.New("ERROR: Invalid link format")
				}
				room1, room2 := parts[0], parts[1]
				//add come and go conections
				if graph.Rooms[room1] == nil {
					continue
				}
				if graph.Rooms[room2] == nil {
					continue
				}
				graph.Rooms[room1].Neighbors = append(graph.Rooms[room1].Neighbors, graph.Rooms[room2])
				graph.Rooms[room2].Neighbors = append(graph.Rooms[room2].Neighbors, graph.Rooms[room1])
				continue

			}

			parts := strings.Fields(line) // Split the line into parts
			if len(parts) != 3 {
				return 0, nil, errors.New("ERROR: invalid room format")
			}

			name := parts[0]
			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])

			if err1 != nil || err2 != nil {
				return 0, nil, errors.New("ERROR: invalid room format")
			}

			if existingRoom, exists := graph.Rooms[name]; exists {
				// If the room name already exists, check the coordinates
				if existingRoom.X != x || existingRoom.Y != y {
					return 0, nil, errors.New("ERROR: duplicate room name with different coordinates")
				}
				// If coordinates are the same, it's a harmless duplicate → skip or continue
				continue
			}

			// Check for different room names with the same coordinates
			coordKey := fmt.Sprintf("%d,%d", x, y)
			if existingRoom, exists := coordinatesMap[coordKey]; exists {
				if existingRoom != name {
					return 0, nil, fmt.Errorf("ERROR: rooms '%s' and '%s' share the same coordinates (%s)", existingRoom, name, coordKey)
				}
			}
			// Create a new Room object
			newRoom := &Room{
				Name: name,
				X:    x,
				Y:    y,
			}

			// Set flags if needed
			if flag {
				StartingRooms++
				newRoom.IsStart = true
				graph.Start = newRoom
				// Reset flag after using it
				FirstFlag = true
				flag = false
			} else if !FirstFlag {
				EndingRooms++
				newRoom.IsEnd = true
				graph.End = newRoom
				FirstFlag = true
			} else {
				ExtraRooms++
			}

			// Now add the new room to the graph map
			graph.Rooms[name] = newRoom
			coordinatesMap[coordKey] = name
		case "start-end":

			if line == "##start" {

				flag = true
				FirstFlag = false
				phase = "rooms"
				continue
			} else if line == "##end" {

				flag = false
				FirstFlag = false
				phase = "rooms"
				continue
			}

		}

	}

	if graph.Start == nil || graph.End == nil { // chack start and end
		return 0, nil, errors.New("ERROR: missing start or end room")
	}
	fmt.Println("Number of Ants", ants)
	fmt.Println("Start Rooms", StartingRooms)
	fmt.Println("Ending Rooms", EndingRooms)
	fmt.Println("Extra Rooms", ExtraRooms)
	fmt.Println("Number of links", tunels)
	fmt.Println(graph.Start.Name)
	fmt.Println(graph.End.Name)

	return ants, graph, nil
}
