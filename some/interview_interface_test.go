package some

import (
	"fmt"
	"io"
	"testing"
)

type coder interface {
	code()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func TestInterface_nil_1(t *testing.T) {
	var c coder
	t.Log(c == nil) // true
	// 动态类型和动态值均为 nil
	// c: <nil>, <nil>
	t.Logf("c: %T, %v\n", c, c)
	
	var g *Gopher
	t.Log(g == nil) // true
	
	c = g
	t.Log(c == nil) // false
	// c 的动态类型为 *interview.Gopher，动态值为 nil
	// c: *interview.Gopher, <nil>
	t.Logf("c: %T, %v\n", c, c)
}

func TestAssert(t *testing.T) {
	x := interface{}(nil)
	t.Logf("Type: %T, value: %v", x, x)
	y := (*int)(nil)
	t.Logf("Type: %T, value: %v", y, y)
	a := x == y   // false
	b := y == nil // true
	// 如果动态类型不存在，则断言总是失败
	_, c := x.(interface{}) // false
	t.Log(a, b, c)
}

type MyErr struct{}

func (m MyErr) Error() string {
	return "My Error"
}

func process() error {
	var err *MyErr = nil
	return err // 这里隐含了类型转换
}

func TestInterface_nil_2(t *testing.T) {
	err := process()
	t.Log(err)        // nil
	t.Log(err == nil) // false
}

/* 如何判断某个类型是否实现了某个接口 */
type myWriter struct{}

func (w myWriter) Write(p []byte) (n int, err error) {
	return
}

func TestJudge(t *testing.T) {
	// 检查 *myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = (*myWriter)(nil)
	// 检查 myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = myWriter{}
}

/* 接口作为函数的入参 */
type S struct{}

func f(x interface{}) {}

func g(x *interface{}) {}

func TestArgumentOfInterface(t *testing.T) {
	s := S{}
	p := &s
	f(s)
	f(p)
	// g(s)  这两种都是错误的用法
	// g(p)
}

/* Stringer 接口*/
type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	// 这里自动调用 c 的 String()，递归调用导致报错！！！
	// return fmt.Sprintf("print: %v", c)
	return ""
}

func TestStringer(t *testing.T) {
	c := &ConfigOne{}
	c.String()
}

/* Any Type */
func TestAnyType(t *testing.T) {
	var x interface{}
	var y interface{} = []int{3, 5,}
	_ = x == x
	_ = x == y
	_ = y == y // 报错，因为两个比较值的动态类型为同一个不可比较类型
}
