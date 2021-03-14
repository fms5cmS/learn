package maze

import "testing"

func TestReadMaze(t *testing.T) {
	maze := ReadMaze("maze.txt")
	for i := range maze {
		t.Log(maze[i])
	}
}
