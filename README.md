# Lem-in: Digital Ant Farm

## Overview

Welcome to Lem-in, a digital version of an ant farm! This project is designed to simulate the movement of ants through a colony of rooms and tunnels, finding the quickest path from the start room to the end room.

## Team Structure
- Stamatis Manousis
- Giorgos Pavrianidis
- Giannis Georgakopoulos
- Edouardos Pavel

## How It Works

Lem-in reads from a file that describes the ants and the colony. The program then finds the quickest path for the ants to travel from the start room to the end room and displays each move the ants make.

## File Structure

The project is organized into several files and directories, each serving a specific purpose:

- **main.go**: The entry point of the application. It reads the input file, validates the graph, finds paths, filters them, and moves the ants.
- **utilities/**: Contains various utility functions and types used throughout the project.
    - **findpaths.go**: Implements the pathfinding algorithm to find all possible paths from the start room to the end room.
    - **filterpaths.go**: Filters overlapping paths to optimize the movement of ants.
    - **input.go**: Parses the input file and constructs the graph representing the colony.
    - **moveants.go**: Simulates the movement of ants through the colony based on the filtered paths.
    - **readinput.go**: Reads the input file name from the command line arguments.
    - **validation.go**: Validates the graph to ensure it has a valid structure and a path from the start room to the end room.
    - **util_test.go**: Contains unit tests for the utility functions.
- **examples/**: Contains example input files that describe different colonies and their configurations.
- **go.mod**: Specifies the Go module and its dependencies.

This structure ensures that the project is modular and easy to navigate, with each component having a clear responsibility.

### Key Features

- **Ant Farm Simulation**: Create a colony with rooms and tunnels.
- **Pathfinding**: Find the quickest way to get ants across the colony.
- **Error Handling**: Handle various invalid or poorly-formatted inputs gracefully.

### Algorithms Used

Lem-in utilizes several algorithms to achieve its functionality:

- **Breadth-First Search (BFS)**: Used to find the shortest path from the start room to the end room. BFS is ideal for this purpose as it explores all possible paths level by level, ensuring the shortest path is found.
- **Depth-First Search (DFS)**: Employed in certain scenarios to explore all possible paths in the colony. DFS is useful for its ability to delve deep into each path, which can be beneficial for specific optimizations and validations.

These algorithms ensure that the ants find the quickest and most efficient paths through the colony.

## Instructions

### Creating the Colony

- **Rooms**: 
    - Rooms must not start with the letter `L` or `#` and must have no spaces.
    - Each room has unique coordinates (integers).
    - Rooms can be linked to multiple other rooms.
    
- **Tunnels**:
    - Tunnels join exactly two rooms.
    - No more than one tunnel can connect the same pair of rooms.
    - Each tunnel can only be used once per turn.

### Ant Movement

- All ants start in the `##start` room and aim to reach the `##end` room.
- Each room can contain only one ant at a time, except for `##start` and `##end`.
- Ants move through tunnels, and each ant can move only once per turn.
- The program displays the ants that moved at each turn.

### Error Handling

- The program checks for various errors, such as:
    - Missing `##start` or `##end` rooms.
    - Invalid room names or coordinates.
    - Duplicate rooms or tunnels.
    - Invalid number of ants.
- Specific error messages are provided for different types of invalid input.

## Example Usage

To run the program, use the following command:
```sh
go run main.go <path_to_file>
```

The program will read the input file, find the quickest path for the ants, and display the moves.

## Bonus
[Here](https://platform.zone01.gr/git/epavel/lem-in-random-map.git) you can find a small side-project that can provide with random maps to test this project with!

## Conclusion

Lem-in is a fun and challenging project that combines pathfinding algorithms with error handling and simulation. We hope you enjoy working on it as much as we did!
