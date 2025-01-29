package utilities

import (
	"os"
	"testing"
)

func TestReadInput(t *testing.T) {
	// Test case 1
	os.Args = []string{"main.go"}
	expected := " "
	actual, _ := ReadInput()
	if actual != expected {
		t.Errorf("Test case 1 failed: expected %s but got %s", expected, actual)
	}

	// Test case 2
	os.Args = []string{"main.go", "input.txt"}
	expected = "input.txt"
	actual, _ = ReadInput()
	if actual != expected {
		t.Errorf("Test case 2 failed: expected %s but got %s", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	// Test case 1
	_, _, err := ParseInput("nonexistentfile.txt")
	if err == nil {
		t.Errorf("Test case 1 failed: expected an error but got nil")
	}

	// Test case 2
	_, _, err = ParseInput("../test.txt")
	if err != nil {
		t.Errorf("Test case 2 failed: expected nil but got %v", err)
	}
}

func TestValidateGraph(t *testing.T) {
	// Test case 1
	graph := &Graph{}
	err := ValidateGraph(graph)
	if err == nil {
		t.Errorf("Test case 1 failed: expected an error but got nil")
	}

	// Test case 2
	graph = &Graph{
		Start: &Room{},
		End:   &Room{},
	}
	err = ValidateGraph(graph)
	if err == nil {
		t.Errorf("Test case 2 failed: expected an error but got nil")
	}

	// Test case 3
	graph = &Graph{
		Start: &Room{},
		End:   &Room{},
		Rooms: map[string]*Room{
			"start": &Room{
				Name: "start",
			},
			"end": &Room{
				Name: "end",
			},
		},
	}
	err = ValidateGraph(graph)
	if err == nil {
		t.Errorf("Test case 3 failed: expected an error but got nil")
	}

	// Test case 4
	// Create rooms
	startRoom := &Room{Name: "start", IsStart: true}
	endRoom := &Room{Name: "end", IsEnd: true}
	roomA := &Room{Name: "A"}
	roomB := &Room{Name: "B"}

	// Set neighbors
	startRoom.Neighbors = []*Room{roomA}
	roomA.Neighbors = []*Room{startRoom, roomB}
	roomB.Neighbors = []*Room{roomA, endRoom}
	endRoom.Neighbors = []*Room{roomB}

	// Create graph
	graph = &Graph{
		Rooms: map[string]*Room{
			"start": startRoom,
			"A":     roomA,
			"B":     roomB,
			"end":   endRoom,
		},
		Start: startRoom,
		End:   endRoom,
	}
	err = ValidateGraph(graph)
	if err != nil {
		t.Errorf("Test case 4 failed: expected nil but got %v", err)
	}
}

func TestFindPaths(t *testing.T) {
	// Create rooms
	startRoom := &Room{Name: "start", IsStart: true}
	endRoom := &Room{Name: "end", IsEnd: true}
	roomA := &Room{Name: "A"}
	roomB := &Room{Name: "B"}

	// Set neighbors
	startRoom.Neighbors = []*Room{roomA}
	roomA.Neighbors = []*Room{startRoom, roomB}
	roomB.Neighbors = []*Room{roomA, endRoom}
	endRoom.Neighbors = []*Room{roomB}

	// Create graph
	graph := &Graph{
		Rooms: map[string]*Room{
			"start": startRoom,
			"A":     roomA,
			"B":     roomB,
			"end":   endRoom,
		},
		Start: startRoom,
		End:   endRoom,
	}

	// Test case 1
	paths := FindPaths(*graph)
	if len(paths) != 1 {
		t.Errorf("Test case 1 failed: expected 1 paths but got %d", len(paths))
	}

	// Test case 2
	// Create rooms
	startRoom = &Room{Name: "start", IsStart: true}
	endRoom = &Room{Name: "end", IsEnd: true}
	roomA = &Room{Name: "A"}
	roomB = &Room{Name: "B"}
	roomC := &Room{Name: "C"}

	// Set neighbors
	startRoom.Neighbors = []*Room{roomA, roomB}
	roomA.Neighbors = []*Room{startRoom, roomC}
	roomB.Neighbors = []*Room{startRoom, roomC}
	roomC.Neighbors = []*Room{roomA, roomB, endRoom}
	endRoom.Neighbors = []*Room{roomC}

	// Create graph
	graph = &Graph{
		Rooms: map[string]*Room{
			"start": startRoom,
			"A":     roomA,
			"B":     roomB,
			"C":     roomC,
			"end":   endRoom,
		},
		Start: startRoom,
		End:   endRoom,
	}

	paths = FindPaths(*graph)
	if len(paths) != 2 {
		t.Errorf("Test case 2 failed: expected 2 paths but got %d", len(paths))
	}
}

func TestFilterPaths(t *testing.T) {
	// Create rooms
	startRoom := &Room{Name: "start", IsStart: true}
	endRoom := &Room{Name: "end", IsEnd: true}
	roomA := &Room{Name: "A"}
	roomB := &Room{Name: "B"}
	roomC := &Room{Name: "C"}

	// Set neighbors
	startRoom.Neighbors = []*Room{roomA, roomB}
	roomA.Neighbors = []*Room{startRoom, roomC}
	roomB.Neighbors = []*Room{startRoom, roomC}
	roomC.Neighbors = []*Room{roomA, roomB, endRoom}
	endRoom.Neighbors = []*Room{roomC}

	// Create graph
	graph := &Graph{
		Rooms: map[string]*Room{
			"start": startRoom,
			"A":     roomA,
			"B":     roomB,
			"C":     roomC,
			"end":   endRoom,
		},
		Start: startRoom,
		End:   endRoom,
	}

	// Test case 1
	paths := FindPaths(*graph)
	filteredPaths := FilterPaths(paths, graph)
	if len(filteredPaths) != 1 {
		t.Errorf("Test case 1 failed: expected 1 paths but got %d", len(filteredPaths))
	}
}

func TestMoveAnts(t *testing.T) {
	// Create rooms
	startRoom := &Room{Name: "start", IsStart: true}
	endRoom := &Room{Name: "end", IsEnd: true}
	roomA := &Room{Name: "A"}
	roomB := &Room{Name: "B"}
	roomC := &Room{Name: "C"}

	// Set neighbors
	startRoom.Neighbors = []*Room{roomA, roomB}
	roomA.Neighbors = []*Room{startRoom, roomC}
	roomB.Neighbors = []*Room{startRoom, roomC}
	roomC.Neighbors = []*Room{roomA, roomB, endRoom}
	endRoom.Neighbors = []*Room{roomC}

	// Create graph
	graph := &Graph{
		Rooms: map[string]*Room{
			"start": startRoom,
			"A":     roomA,
			"B":     roomB,
			"C":     roomC,
			"end":   endRoom,
		},
		Start: startRoom,
		End:   endRoom,
	}

	// Test case 1
	paths := FindPaths(*graph)
	antsCount := 5
	MoveAnts(paths, antsCount, graph)
}
