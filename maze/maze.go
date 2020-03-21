package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

var directions = [4]point{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
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

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(maze [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(maze) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(maze[p.i]) {
		return 0, false
	}
	return maze[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, d := range directions {
			next := cur.add(d)
			// maze at next is 0 not 1 and not out of bound
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// and steps at next is 0 (i.e. not visited)
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// and next != start
			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}
