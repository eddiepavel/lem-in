package utilities

import (
	"errors"
)

func ValidateGraph(graph *Graph) error {
	if graph.Start == nil {
		return errors.New("graph is missing a start room")
	}
	if graph.End == nil {
		return errors.New("graph is missing a start room")
	}
	// namesMap := make(map[string][2]int)
	// coordinatesMap := make(map[string]string)

	// for _, room := range graph.Rooms {
	// 	coordKey := fmt.Sprintf("%d,%d", room.X, room.Y)
	// 	// Check for duplicate room names with different coordinates
	// 	if coords, exists := namesMap[room.Name]; exists {
	// 		if coords != [2]int{room.X, room.Y} {
	// 			return errors.New("duplicate room name with different coordinates")
	// 		}
	// 	}
	// 	// Check for different room names with the same coordinates
	// 	if existingRoom, exists := coordinatesMap[coordKey]; exists {
	// 		if existingRoom != room.Name {
	// 			return errors.New("different room names share the same coordinates")
	// 		}
	// 	}
	// 	namesMap[room.Name] = [2]int{room.X, room.Y}
	// 	coordinatesMap[coordKey] = room.Name
	// 	fmt.Printf("Checking room: %s at (%d, %d)\n", room.Name, room.X, room.Y)
	// }
	return nil
}
