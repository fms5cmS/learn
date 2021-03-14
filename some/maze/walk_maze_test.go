package maze

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	maze := ReadMaze("maze.txt")
	start, end := Point{0, 0}, Point{5, 6}
	steps := Walk(maze, start, end)
	for _, row := range steps {
		for _, v := range row {
			// 3位对齐
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
}
