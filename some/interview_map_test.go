package some

import "testing"

func TestCreateMap(t *testing.T) {
	var slice []int
	// slice[0] = 1  // index out of range [0] with length 0
	slice = append(slice,1)
	t.Log(slice)  // 【1】
	var m1 map[string]int
	t.Log(m1 == nil)  // true
	// m1["s"] = 3  m1 还未分配内存，无法直接赋值，这里报错
	// 字面量的方式初始化
	m2 := map[string]int{"id": 2}
	t.Log(m2)
	// 推荐使用 make 来初始化 map，虽然可以指定容量，但会被忽略
	m3 := make(map[string]int)
	t.Log(m1, m3) // map[] map[]
}
