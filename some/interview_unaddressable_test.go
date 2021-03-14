package some

import "testing"

// 不可变的
// func TestImmutable(t *testing.Test) {
// 	// 1. 常量是不可寻址的，编译报错
// 	const a = 100
// 	t.Log(&a)
// 	// 2. 基本类型字面量不可寻址，编译报错
// 	t.Log(&12)
// 	// 3. 字符串类型是采用只读(不可变)的 byte slice 存储的，所以基于它的索引或切片的结果值也都是不可寻址的
// 	str := "abc"
// 	t.Log(&str[0])    // 基于字符串类型的索引的结果值是不可寻址的，编译报错
// 	t.Log(&str[0:1])  //基于字符串类型的切片的结果值是不可寻址的，编译报错
// 	s := str[0]
// 	t.Log(&s)   // 这样是合法的
// 	// 4. 函数是不可变的，且拿到指向一段代码的指针是不安全的
// 	_ = &func(x,y int) int {  // 字面量代表的函数不可寻址
// 		return x+y
// 	}
// 	_ = &(fmt.Sprintf) // 标识符代表的函数不可寻址
// 	_ = &(fmt.Sprintln("abc")) // 对函数的调用结果值不可寻址
// }

// func TestTemporaryResult(t *testing.Test) {  // 临时结果
// 	// 1. 算术操作的结果值不可寻址
// 	t.Log(&(12+13))
// 	a,b := 1,2
// 	t.Log(&(a+b))  // 也是不可寻址的
// 	c := a+b
// 	t.Log(&c)  // 将临时结果赋值给变量后就是可寻址的了
// 	// 2. 数组值、切片值或字典值的字面量的表达式会产生临时结果，所以是不可寻址的
// 	t.Log(&([3]int{1, 2, 3}[0]))
// 	// 如果针对的是数组类型或切片类型的变量，那么索引或切片的结果值就都不属于临时结果了，是可寻址的
// }

type Test struct {
	ls []int
}

func foo_T(t Test){
	t.ls[0] = 100
}

func TestStructField(t *testing.T) {
	test := Test{[]int{1, 2, 3}}
	foo_T(test)
	t.Log(test.ls) // 100 2 3
}

type x struct {}

func (x *x) test() {
	println(x)
}

func getX() x {
	return x{}
}

func TestTmp(t *testing.T) {
	var a *x
	a.test()
	// x{}.test()  // x{} 是不可寻址的
	b := x{}
	b.test()
	// getX() 直接返回的 x{} 是不可寻址的；不可寻址的结构体不能调用带结构体指针接收者的方法
	// getX().test()
}