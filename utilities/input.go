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
	Count       int
	Ants        Ant
	AllPaths    [][]string
	FinalPaths  [][]string
	FinalSteps  int
	FinalOutput string
	Rooms       map[string]*Room
	Start       *Room
	End         *Room
}

func NewGraph() *Graph {
	return &Graph{
		Rooms: make(map[string]*Room),
	}
}

func (g *Graph) ParseInput() error {
	file, err := ReadInput()
	if err != nil {
		return err
	}
	var startRoom, endRoom, anyRoom *Room

	f, err := os.Open(file)
	if err != nil {
		return err
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
						return errors.New("ERROR: multiple ##start markers")
					}
					startRoom = &Room{}
					phase = "start"
				} else if line == "##end" {
					if endRoom != nil {
						return errors.New("ERROR: multiple ##end markers")
					}
					endRoom = &Room{}
					phase = "end"
				} else {
					return errors.New("ERROR: invalid double # marker")
				}
			}
			continue
		}

		switch phase {
		case "ants":
			g.Count, err = parseAnts(line)
			if err != nil {
				return err
			}
			phase = "rooms"
		case "rooms":
			if strings.Contains(line, "-") {
				err = g.parseLink(line)
			} else {
				err = g.parseRoom(line, anyRoom)
			}
			if err != nil {
				return err
			}
		case "start":
			err = g.parseRoom(line, startRoom)
			if err != nil {
				return err
			}
			startRoom.IsStart = true
			g.Start = startRoom
			phase = "rooms"
		case "end":
			err = g.parseRoom(line, endRoom)
			if err != nil {
				return err
			}
			endRoom.IsEnd = true
			g.End = endRoom
			phase = "rooms"
		}
	}

	if g.Start == nil || g.End == nil {
		return errors.New("ERROR: missing start or end room")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseAnts(line string) (int, error) {
	ants, err := strconv.Atoi(line)
	if err != nil || ants <= 0 {
		return 0, errors.New("error: invalid number of ants")
	}
	return ants, nil
}

func (g *Graph) parseRoom(line string, newRoom *Room) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return errors.New("ERROR: invalid room format")
	}

	name := parts[0]
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "l") {
		return errors.New("error: Room name cannot start from L or l")
	}

	x, err1 := strconv.Atoi(parts[1])
	y, err2 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil {
		return errors.New("ERROR: invalid room format")
	}

	if existingRoom, exists := g.Rooms[name]; exists {
		if existingRoom.X != x || existingRoom.Y != y {
			return errors.New("ERROR: duplicate room name with different coordinates")
		}
		return nil
	}

	coordKey := fmt.Sprintf("%d,%d", x, y)
	for _, room := range g.Rooms {
		if room.X == x && room.Y == y && room.Name != name {
			return fmt.Errorf("ERROR: rooms '%s' and '%s' share the same coordinates (%s)", room.Name, name, coordKey)
		}
	}

	if newRoom == nil {
		newRoom = &Room{}
	}

	newRoom.X, newRoom.Y, newRoom.Name = x, y, name
	g.Rooms[name] = newRoom

	return nil
}

func (g *Graph) parseLink(line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return errors.New("ERROR: Invalid link format")
	}

	room1, room2 := parts[0], parts[1]
	if g.Rooms[room1] == nil || g.Rooms[room2] == nil {
		return nil
	}

	g.Rooms[room1].Neighbors = append(g.Rooms[room1].Neighbors, g.Rooms[room2])
	g.Rooms[room2].Neighbors = append(g.Rooms[room2].Neighbors, g.Rooms[room1])

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
