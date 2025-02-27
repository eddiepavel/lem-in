package main

import (
	"fmt"
	"main/utilities"
)

func main() {
	// validate input
	ants, graph, err := utilities.ParseInput()
	if err != nil {
		fmt.Println("Parsing failed: ", err)
		return
	}
	// validate graph
	if err := utilities.ValidateGraph(graph); err != nil {
		fmt.Println("Validation failed:", err)
		return // Stop execution if validation fails
	}
	// print input
	utilities.Print()
	// all valid paths available but with ovelapping
	paths := utilities.FindPaths(*graph)
	// filter overlapping paths
	filterPaths := utilities.FilterPaths(paths, graph, &ants)
	// move ants
	output, _ := utilities.MoveAnts(filterPaths, ants, graph)
	fmt.Println(output)
}
