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

	fmt.Println(ants, graph)
}
