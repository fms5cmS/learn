package maze

func PrintPath(steps [][]int, end, start Point) [][]int {
	path := make([][]int, len(steps))
	for i := range path {
		path[i] = make([]int, len(steps[i]))
	}
	val, _ := end.Val(steps)
	path[end.I][end.J] = val
	queue := []Point{end}
flag:
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curVal, _ := cur.Val(steps)
		for _, dir := range dirs {
			next := cur.Move(dir)
			if next == start {
				break flag
			}
			val, _ := next.Val(steps)
			if val == curVal-1 {
				path[next.I][next.J] = val
				queue = append(queue, next)
			} else {
				continue
			}
		}
	}
	return path
}
