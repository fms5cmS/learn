package maze

type Point struct {
	I, J int
}

// 这里使用的不是指针接收者，并返回一个 Point
func (p Point) Move(r Point) Point {
	return Point{
		I: p.I + r.I,
		J: p.J + r.J,
	}
}

// 确定在地图 grid 中坐标 Point 的值，以及该坐标是否越界
// 越界返回 true
func (p Point) Val(grid [][]int) (v int, out bool) {
	if p.I < 0 || p.I >= len(grid) {
		return 0, true
	}
	if p.J < 0 || p.J >= len(grid[p.I]) {
		return 0, true
	}
	return grid[p.I][p.J], false
}
