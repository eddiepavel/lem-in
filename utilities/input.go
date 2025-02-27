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

func ParseInput() (int, *Graph, error) {
	file, err := ReadInput()
	if err != nil {
		return 0, nil, err
	}

	graph := &Graph{Rooms: make(map[string]*Room)}
	var ants int
	var startRoom, endRoom *Room

	f, err := os.Open(file)
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	phase := "ants"
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			if strings.HasPrefix(line, "##") {
				if line == "##start" {
					if startRoom != nil {
						return 0, nil, errors.New("ERROR: multiple ##start markers")
					}
					phase = "start"
				} else if line == "##end" {
					if endRoom != nil {
						return 0, nil, errors.New("ERROR: multiple ##end markers")
					}
					phase = "end"
				} else {
					return 0, nil, errors.New("ERROR: invalid double # marker")
				}
			}
			continue
		}

		switch phase {
		case "ants":
			ants, err = parseAnts(line)
			if err != nil {
				return 0, nil, err
			}
			phase = "rooms"
		case "rooms":
			if strings.Contains(line, "-") {
				err = parseLink(line, graph)
			} else {
				_, err = parseRoom(line, graph)
			}
			if err != nil {
				return 0, nil, err
			}
		case "start":
			startRoom, err = parseRoom(line, graph)
			if err != nil {
				return 0, nil, err
			}
			startRoom.IsStart = true
			graph.Start = startRoom
			phase = "rooms"
		case "end":
			endRoom, err = parseRoom(line, graph)
			if err != nil {
				return 0, nil, err
			}
			endRoom.IsEnd = true
			graph.End = endRoom
			phase = "rooms"
		}
	}

	if graph.Start == nil || graph.End == nil {
		return 0, nil, errors.New("ERROR: missing start or end room")
	}

	if err := scanner.Err(); err != nil {
		return 0, nil, err
	}

	return ants, graph, nil
}

func parseAnts(line string) (int, error) {
	ants, err := strconv.Atoi(line)
	if err != nil || ants <= 0 {
		return 0, errors.New("error: invalid number of ants")
	}
	return ants, nil
}

func parseRoom(line string, graph *Graph) (*Room, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return nil, errors.New("ERROR: invalid room format")
	}

	name := parts[0]
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "l") {
		return nil, errors.New("error: Room name cannot start from L or l")
	}

	x, err1 := strconv.Atoi(parts[1])
	y, err2 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil {
		return nil, errors.New("ERROR: invalid room format")
	}

	if existingRoom, exists := graph.Rooms[name]; exists {
		if existingRoom.X != x || existingRoom.Y != y {
			return nil, errors.New("ERROR: duplicate room name with different coordinates")
		}
		return existingRoom, nil
	}

	coordKey := fmt.Sprintf("%d,%d", x, y)
	for _, room := range graph.Rooms {
		if room.X == x && room.Y == y && room.Name != name {
			return nil, fmt.Errorf("ERROR: rooms '%s' and '%s' share the same coordinates (%s)", room.Name, name, coordKey)
		}
	}

	newRoom := &Room{Name: name, X: x, Y: y}
	graph.Rooms[name] = newRoom

	return newRoom, nil
}

func parseLink(line string, graph *Graph) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return errors.New("ERROR: Invalid link format")
	}

	room1, room2 := parts[0], parts[1]
	if graph.Rooms[room1] == nil || graph.Rooms[room2] == nil {
		return nil
	}

	graph.Rooms[room1].Neighbors = append(graph.Rooms[room1].Neighbors, graph.Rooms[room2])
	graph.Rooms[room2].Neighbors = append(graph.Rooms[room2].Neighbors, graph.Rooms[room1])

	return nil
}

func Print() {
	f, _ := os.Open(os.Args[1])
	f.Seek(0, 0) // Reset the file pointer to the beginning
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println()
}

func ReadInput() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("usage: go run main.go <input_file_name>")
	}
	return os.Args[1], nil
}
