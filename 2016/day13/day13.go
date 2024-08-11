package main

import (
	"fmt"
)

func main() {
	test()
	part1()
	part2()
}

func test() {
	maze := Maze{}
	maze.init(10, 7, 10)
	solved, steps := maze.solve(&maze.grid[1][1], &maze.grid[4][7])
	if solved {
		fmt.Println("A solution was found. Optimal path is", steps, "steps")
		maze.print()
	}
}

func part1() {
	maze := Maze{}
	maze.init(50, 50, 1352)
	solved, steps := maze.solve(&maze.grid[1][1], &maze.grid[39][31])
	if solved {
		fmt.Println("A solution was found. Optimal path is", steps, "steps")
		maze.print()
	}
}

func part2() {
	maze := Maze{}
	maze.init(50, 50, 1352)
	locations := 0

	for y := range maze.height {
		for x := range maze.width {
			if maze.isWall(x, y) == false {
				solved, steps := maze.solve(&maze.grid[1][1], &maze.grid[y][x])
				maze.clear()

				if solved && steps <= 50 {
					locations++
				}
			}
		}
	}

	fmt.Println(locations)
}
