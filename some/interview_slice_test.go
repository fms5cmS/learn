package some

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(arr[:6]) // 0, 1, 2, 3, 4, 5, _, _, _, _
	t.Log(arr[2:]) //      2, 3, 4, 5, 6, 7, 8, 9
	t.Log(arr[:])  // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	
	s1 := arr[2:5]  //      2, 3, 4, _, _, _, _, _  容量为 8
	s2 := s1[2:6:7] //           4, 5, 6, 7, _,    容量为 5
	s2 = append(s2, 100)
	s2 = append(s2, 200) // 容量不足，扩容导致底层数组发生改变！！！
	s1[2] = 20
	t.Log(s1)  //      2, 3,20, _, _, _, _, _
	t.Log(s2)  //            4, 5, 6, 7,100,200
	t.Log(arr) // 0, 1, 2, 3,20, 5, 6, 7,100, 9
}

// append 的第一个参数必须是切片，即使是切片指针也不行，append 的返回值相同
func TestSliceAppend_first_arg(t *testing.T) {
	list := *new([]int)
	list = append(list, 1)
	println(list)
}

// 切片的底层数据结构是结构体，append 函数修改了切片的底层数据结构！
// 在 testSliceValue 中传入的是切片结构体的副本，切片本身不会被改变
// 在 testSlicePoint 中传入指针的副本，所以切片本身也发生了改变
func TestSlice_Struct(t *testing.T) {
	testSliceValue := func(stack []string) {
		stack = append(stack, "e")
	}
	testSlicePoint := func(stack *[]string) {
		*stack = append(*stack, "f")
	}
	arr := [5]string{"a", "b", "c", "d"}
	slice := arr[:4]
	testSliceValue(slice)
	t.Log(arr)   // a b c d e
	t.Log(slice) // a b c d 注意，尽管底层数组发生了改变，但 slice 并未改变
	testSlicePoint(&slice)
	t.Log(arr)   // a b c d f
	t.Log(slice) // a b c d f 注意，底层数组和 slice 都发生了改变
}

func TestSlice_Arg(t *testing.T) {
	change := func(s ...int) {
		s = append(s, 3)
	}
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	t.Log(slice) // [1 2 0 0 0]
	change(slice[0:2]...)
	t.Log(slice) // [1 2 3 0 0]
}

func TestTwoDimensionalSlice(t *testing.T) {
	file, _ := os.Open("maze.txt")
	row, col := 3, 4
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			// 注意，文件中的换行符必须换成 LF，不能使用 Windows 默认的换行符
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	t.Log(maze)
}

func TestDeleteElementFromSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	// 删除索引为 3 的元素
	s = append(s[:3], s[4:]...)
	// 删除中间元素，len-1、cap 不变
	// [1 2 3 5 6], len = 5,cap = 6
	t.Logf("%v, len = %d,cap = %d", s, len(s), cap(s))
	// 删除第一个元素，len-1、cap-1
	// [2 3 5 6], len = 4,cap = 5
	s = s[1:]
	t.Logf("%v, len = %d,cap = %d", s, len(s), cap(s))
	// 删除最后一个元素，len-1、cap 不变
	// [2 3 5], len = 3,cap = 5
	s = s[:len(s)-1]
	t.Logf("%v, len = %d,cap = %d", s, len(s), cap(s))
}

func TestGet(t *testing.T) {
	x := make([]int, 2, 10)
	_ = x[6:10]
	_ = x[6:] // 报错，x 的长度为 2，这里的起始索引 > 2
	_ = x[2:]
}

func TestArrRange(t *testing.T) {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6} // 数组
	maxIndex2 := len(numbers2) - 1         // 5
	// 传入的是数组的副本，所以遍历的是 [1, 2, 3, 4, 5, 6] 这些数字
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	t.Log(numbers2) // [7 3 5 7 9 11]
}

func TestSliceRange(t *testing.T) {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	// 传入的是切片的副本，但由于下面的每次修改都对底层数据进行了修改，所以切片副本的底层数据也在不断改变
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	t.Log(numbers2) // [22 3 6 10 15 21]
}
