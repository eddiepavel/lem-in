package main

import (
	"fmt"
	"main/utilities"
)

// place your progress in a file inside the utilities dir
// step-by-step process
// ||||||||||||||||||||
// \/\/\/\/\/\/\/\/\/\/

func main() {
	// file input + validation (giorgos)
	// map validation (stamatis)
	// map desing (giannis)
	// recursive path finding (eddie)
	// filter non-overlapping paths (stamatis)
	// brute force and calculation of max steps (eddie)
	// deploy ants (giannis)
	// file output (giorgos)
	file, err := utilities.ReadInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	ants, graph, err := utilities.ParseInput(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Running graph validation...")
	if err := utilities.ValidateGraph(graph); err != nil {
		fmt.Println("Validation failed:", err)
		return // Stop execution if validation fails
	}

	//all valid paths available but with ovelapping
	paths := utilities.FindPaths(*graph)

	fmt.Println("Filtering paths...")
	//filter overlapping paths
	filterPaths := utilities.FilterPaths(paths, graph)

	for _, path := range filterPaths {
		fmt.Println(path)
	}

	fmt.Printf("%d %+v\n", ants, graph)
}
