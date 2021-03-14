package some

import (
	"fmt"
	"testing"
)

// for range 使用短变量声明(:=)的形式迭代变量时，变量 i、value 在每次循环体中都会被重用，而不是重新声明!!
func TestRange_ShortVar(t *testing.T) {
	type foo struct {
		bar string
	}
	s1 := []foo{{"A"}, {"B"}, {"C"}}
	s2 := make([]*foo, len(s1))
	// s2 每次填充的都是临时变量 value 的地址，而在最后一次循环中，value 被赋值为{c}。
	// 因此，s2 输出的时候显示出了三个 &{C}。
	for i, v := range s1 {
		s2[i] = &v
	}
	t.Log(s1)
	t.Log(s2[0], s2[1], s2[2])
	// 可修改为以下：
	for i := range s1 {
		s2[i] = &s1[i]
	}
	t.Log(s2[0], s2[1], s2[2])
}

func TestRange(t *testing.T) {
	var k = 9
	for k = range []int{} {
	}
	t.Log(k) // 9
	for k = 0; k < 3; k++ {
	}
	t.Log(k) // 3
	for k = range (*[3]int)(nil) {
	}
	t.Log(k) // 2，数组的索引
}

type TN struct {
	n int
}

func TestRangeArr(t *testing.T) {
	ts := [2]TN{}
	for i,v := range ts{
		switch i {
		case 0:
			v.n = 3  // range 使用的是 ts 的副本，这里的赋值操作不会影响原数组
			ts[1].n = 9
		case 1:
			fmt.Print(v.n," ")  // 0
		}
	}
	fmt.Println(ts)  // [{0} {9}]
	
	for i, v := range &ts {
		switch i {
		case 0:
			v.n = 3  // 尽管 range 使用的是 ts 的指针，但数组元素是结构体，v 为原数组元素的副本，对副本的修改不影响原数组元素
			ts[1].n = 9
		case 1:
			fmt.Print(v.n," ")  // 9
		}
	}
	fmt.Println(ts)  // [{0} {9}]
}
