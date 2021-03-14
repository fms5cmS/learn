package some

import "testing"

type t_slice []int

// Named Type(具名类型)：内置类型、使用 `type` 关键字声明的类型
// Unnamed Type：基于已有的 Named Type 组合在一起的类型，如 `[]string`、`map[string]int`、`interface{}`、`struct{}`
// 如果两个变量，其底层类型相同，且至少有一个是 Unnamed Type，则这两个变量是可以直接转换的！
func TestTransform(t *testing.T) {
	a := []int{1, 2, 3, 4}
	var b t_slice = a
	t.Log(b)
	
	var x t_slice = make(t_slice, 2)
	x = append(x, 10)
	x = append(x, 20)
	var y []int = x
	t.Log(y)
}

type fo struct {
	Val int
}

type ba struct {
	Val int
}

func TestStructTransform(t *testing.T) {
	a := fo{3,}
	b := ba{3,}
	// 注意，两个结构体的字段名也必须一致
	t.Log(a == fo(b)) // true
}
