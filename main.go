package main

import (
	"fmt"
	"main/utilities"
)

func main() {
	// validate input
	graph := &utilities.Graph{}
	err := graph.ParseInput()
	if err != nil {
		fmt.Println("Parsing failed: ", err)
		return
	}
	// validate graph
	if err = graph.ValidateGraph(); err != nil {
		fmt.Println("Validation failed:", err)
		return // Stop execution if validation fails
	}
	// print input
	utilities.Print()
	// all valid paths available but with ovelapping
	graph.FindPaths()
	// filter overlapping paths
	graph.FilterPaths()
	// move ants
	output, _ := utilities.MoveAnts(filterPaths, ants, graph)
	fmt.Println(output)
}
