package some

import (
	"bytes"
	"sync"
	"testing"
)

type syncBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func TestBuiltinFunc(t *testing.T) {
	// make() 仅用于 slice、map、channnel 类型的内存分配和相应内部结构的初始化，返回 T 类型（不是 \*T）的值（不是零值）。
	// new() 用于分配内存，并返回该类型(字段值均为字段类型的零值)的指针
	test_new := new(syncBuffer)
	var test_var syncBuffer
	t.Logf("the type of new: %T; the type of var: %T", test_new, test_var)
}

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func TestAnonymousFunc(t *testing.T) {
	f := F(5)
	defer func() {
		println(f()) // 8
	}()
	defer println(f()) // 6
	i := f()
	println(i) // 7
}

// 该函数返回一个闭包
func appendStr() func(string) string {
	t := "Hello"
	// 使用了自由变量 t，所以是一个闭包
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

// 闭包捕获的变量和常量是引用传递而不是值传递！
func TestClosure(t *testing.T) {
	// a、b 都是闭包，绑定了各自的 t 值
	a := appendStr()
	b := appendStr()
	t.Log(a("A")) // Hello A
	t.Log(b("B")) // Hello B
	
	t.Log(a("AA")) // Hello A AA
	t.Log(b("BB")) // Hello B BB
}

// 可变函数是指针传递
func TestVarFunc(t *testing.T) {
	hello := func(num ...int) { // 本质上 ...int 是一个切片
		num[0] = 18
	}
	i := []int{5, 6, 7}
	hello(i...)
	t.Log(i[0]) // 18
}
