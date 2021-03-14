package some

import (
	"fmt"
	"testing"
)

// 多个 defer 在执行时 LIFO
// 输出结果为：2  1 panic: error occurred
func TestDeferStack(t *testing.T) {
	defer println(1)
	defer println(2)
	panic("error occurred")
	println(3)
}

// defer 的参数
func TestArgumentOfDeferFunc(t *testing.T) {
	calc := func(index string, a, b int) int {
		ret := a + b
		println(index, a, b, ret)
		return ret
	}
	a := 1
	b := 2
	// 参数在执行 defer 时就会被赋值，而不是这个调用执行时赋值，所以会先计算 calc("10", a, b)
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// defer 函数用到的接收者
type slice []int

func (s *slice) Add(elem int) *slice {
	*s = append(*s, elem)
	fmt.Println(elem)
	return s
}

func NewSlice() slice {
	return make(slice, 0) // 注意这里可以使用 make 来创建 []int 重定义后的类型
}

func TestArgument_receiver(t *testing.T) {
	s := NewSlice()
	// s.Add(1) 会先于 s.Add(3) 执行
	defer s.Add(1).Add(2)
	s.Add(3)
}

// 三步拆解：
//  1. r = n + 1 = 4
//  2. 执行第二个 defer，但此时 f() 未定义，故引发异常
//  	 执行第一个 defer，r = r + n = 7，recover 异常，程序正常执行
//  3. return
// 注意：如果没有 recover() 程序会产生 panic，f(3) 报错无法返回值，程序只输出报错信息
func TestDeferWithPanic(t *testing.T) {
	t.Log(f_for_defer(3))
}

func f_for_defer(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()
	var f func()
	defer f()
	f = func() {
		r += 2
	}
	println("test...") // 可以正常打印
	return n + 1
}

type Person struct {
	age int
}

// defer 函数定义时对外部变量的引用
// 3 3 20
func TestDefer1(t *testing.T) {
	person := &Person{20}
	// 传入 20
	defer t.Log(person.age)
	// 外部变量作为函数参数，最后参数 person.age 被修改为 3
	defer func(p *Person) {
		t.Log(p.age)
	}(person)
	// 外部变量作为闭包引用，defer 真正调用时根据整个上下文来确定值
	defer func() {
		t.Log(person.age)
	}()
	// 这里仅修改了引用对象的成员值
	person.age = 3
}

// 3 20 20
func TestDefer2(t *testing.T) {
	person := &Person{20}
	defer t.Log(person.age)
	defer func(p *Person) {
		t.Log(p.age)
	}(person)
	defer func() {
		t.Log(person.age)
	}()
	// 注意：这里修改了 person 的引用！！！！
	person = &Person{3}
}

var a = true

// 2 1
func TestDefer3(t *testing.T) {
	defer func() {
		println("1")
	}()
	if a == true {
		println("2")
		return
	}
	// 这里的函数尚未注册就已经结束了
	defer func() {
		println("3")
	}()
}
