package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{0, 1}, {-1, 0}, {0, -1}, {1, 0},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grad [][]int) (int, bool) {
	if p.i < 0 || p.i > len(grad)-1 {
		return 0, false
	}

	if p.j < 0 || p.j > len(grad[p.i])-1 {
		return 0, false
	}

	return grad[p.i][p.j], true
}

func workMaze(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for i := range dirs {
			next := current.add(dirs[i])

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			queue = append(queue, next)
			curSteps, _ := current.at(steps)
			steps[next.i][next.j] = curSteps + 1
		}
	}

	return steps
}

func main() {
	maze := readMaze("gobasic/maze/maze.txt")

	steps := workMaze(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}

}
