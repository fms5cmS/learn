package some

import (
	"math"
	"testing"
)

func TestConst(t *testing.T) {
	// 1. 指定常量类型
	const filename string = "abc.txt"
	// 2. 没有指定 a、b 的类型，故可以将 a、b 作为任意类型使用
	const a, b = 3, 4
	t.Logf("%T", a)
	c := math.Sqrt(a*a + b*b) // 可以正确编译，注： Sqrt 的参数类型为 float64
	t.Logf("%T", c)        // float64
	// 3. 可以给常量赋一个编译期运算的值
	const mask = 1 << 3
	// 4. math.Sin() 的调用是在运行时发生的，所以下面的语句会报错
	// const sin = math.Sin(math.Pi/4)
	
	// 5. 常量组中如不指定类型和初始化值，则与上一行非空常量相同！！
	const (
		x uint16 = 120
		y
		s = "abc"
		z
	)
	// the type of x、y、s、z: uint16, uint16, string, string
	t.Logf("the type of x、y、s、z: %T, %T, %T, %T", x, y, s, z)
	// the vale of x、y、s、z: 120, 120, abc, abc
	t.Logf("the vale of x、y、s、z: %v, %v, %v, %v", x, y, s, z)
}
