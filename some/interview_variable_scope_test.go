package some

import "testing"

// 变量隐藏。
// 使用变量简短声明符号 := 时，如果符号左边有多个变量，只需要保证至少有一个变量是新声明的，并对已定义的变量尽进行赋值操作。
// 但如果出现作用域之后，就会导致变量隐藏的问题
func TestVariableHide(t *testing.T) {
	x := 1
	t.Log(x) // 1
	{
		t.Log(x) // 1
		i, x := 2, 2
		t.Log(i, x) // 2 2
	}
	t.Log(x) // 1
}

func TestVariableHideWithReturn(t *testing.T) {
	test := func(i int) (ret int) {
		ret = i * 2
		if ret > 10 {
			// ret := 10   // 这里编译报错，ret is shadowed during return
			return
		}
		return
	}
	t.Log(test(10))
}

//
var f_v = func(i int) {
	print("x")
}

func TestVariable(t *testing.T) {
	f_v := func(i int) {
		print(i)
		if i > 0 {
			f_v(i - 1) // 这里调用的是已经完成的全局变量的那个 f_v() 函数
		}
	}
	f_v(10) // 输出 10x
}

//
var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	println(*p)
}

// := 定义的变量，如果新变量与同名已定义的变量不在同一作用域，则会新定义这个变量
// 所以这里的 p 与全局变量的 p 不是同一个变量，全局变量的 p==nil，在调用 bar() 时会报错：
// panic: runtime error: invalid memory address or nil pointer dereference
func TestVariableScope_fatal(t *testing.T) {
	p, err := foo()
	if err != nil {
		println(err)
		return
	}
	bar()
	println(*p)
}

func TestVariableScope_right(t *testing.T) {
	var err error
	p, err = foo()
	if err != nil {
		println(err)
		return
	}
	bar()
	println(*p)
}
