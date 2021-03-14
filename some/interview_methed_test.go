package some

import (
	"fmt"
	"testing"
)
// 类型再定义
type myString string

// func (ms *myString) String() string {
// 	return fmt.Sprintf("the type myString is %T", ms)
// }

// 这里对重定义的类型 myString 实现了 Stringer 接口
// 注意 return 语句不能写 fmt.Sprint(s)！会递归调用 String() 方法从而报错
// return 语句可以写成 fmt.Sprint(string(s))，因为 string 类型没有 String() 方法
func (ms myString) String() string {
	return fmt.Sprintf("the type myString is %T", ms)
}

func TestMethod(t *testing.T) {
	var ms myString = "666"
	t.Log(ms)  // the type myString is some.myString
}

type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v: %p, %v\n", n, *n)
}

func (n N) test() {
	println(n)
}

func (n *N) testPointer() {
	println(*n)
}

// https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648467188&idx=2&sn=1a4dc21485abc97a5cba01079c8ece37&chksm=f247409bc530c98dbd06fa23b8aad3e85bc6f2d046452fac075ca6c5a7f065b8931635926050&scene=21#wechat_redirect
func TestMethodExpression_struct(t *testing.T) {
	var n N = 10
	
	// p := &n
	// 不能使用多级指针调用方法，编译报错
	// p1 := &p
	// p1.value()
	// p1.pointer()
	
	n++
	func1 := N.test    // 通过类型引用的方法表达式被还原成普通函数样式
	func1(n)        // 类型 N 的方法集中有 test，故可以调用，输出 11
	// N.test(n)  也可以直接使用方法表达式调用
	
	n++
	func2 := (*N).test
	func2(&n)      // 类型 *N 的方法集中有 test，可以调用。输出 12
	// (*N).test(n)
}

func TestMethodExpression_instance(t *testing.T) {
	var n N = 10
	p:=&n
	n++
	func1 := n.test
	n++
	func2 := p.test
	n++
	t.Log(n)  // 13
	func1()   // 11
	func2()   // 12
}

func TestMethodExpression_pointer(t *testing.T) {
	var n N = 10
	p := &n
	n++
	func1 := n.testPointer
	n++
	func2 := p.testPointer
	n++
	t.Log(n)  // 13
	func1()   // 13
	func2()   // 13
}