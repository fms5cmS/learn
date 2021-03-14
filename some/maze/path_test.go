package maze

import (
	"fmt"
	"testing"
)

func TestPrintPath(t *testing.T) {
	maze := ReadMaze("maze.txt")
	start, end := Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1}
	steps := Walk(maze, start, end)
	fmt.Println("steps: ")
	for _, row := range steps {
		for _, v := range row {
			// 3位对齐
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
	fmt.Println("path:")
	path := PrintPath(steps, end, start)
	for _, row := range path {
		for _, v := range row {
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
}
