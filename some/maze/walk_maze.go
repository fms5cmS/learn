package maze

// 四个方向：上、左、下、右
var dirs = [4]Point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func Walk(maze [][]int, start, end Point) [][]int {
	// steps 对应于 maze，记录从 start 走到对应迷宫当前坐标的步数
	// maze 记录迷宫情况
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 即将遍历的 Point 队列，初始值只有起点是已到达但未探索的点
	queue := []Point{start}
flag:
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 以 cur 开始向四个方向探索
		for _, dir := range dirs {
			next := cur.Move(dir)
			
			// 判断 next 这个坐标在迷宫中是否越界或遇到墙(1)
			v, out := next.Val(maze)
			if out || v == 1 {
				continue // 跳过这个 next，探索另一个方向
			}
			// 判断 next 这个坐标是否已经走过
			// 这里不用判断越界，因为 steps 和 maze 的行、列是一致的，maze 中如果没有越界，steps 也不会越界
			v, _ = next.Val(steps)
			if v != 0 {
				continue
			}
			// 走回到了起点
			if next == start {
				continue
			}
			
			// 向 steps 中填入值
			curStep, _ := cur.Val(steps)
			steps[next.I][next.J] = curStep + 1
			
			if next == end {
				break flag
			}
			// 将新发现的点加入队列
			queue = append(queue, next)
		}
	}
	return steps
}
