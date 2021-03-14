package some

import "testing"

// Go 中的取反操作符为 ^，其他语言通常为 ~
// 对有符号整数 a 取反，结果为 -(a+1)
// 对于无符号整数来说就是按位取反
func TestNegate(t *testing.T) {
	var a int8 = 3
	t.Logf("^%b=%b %d", a, ^a, ^a)
	var b uint8 = 3
	t.Logf("^%b=%b %d", b, ^b, ^b)
	var c int8 = -3
	t.Logf("^(%b)=%b %d", c, ^c, ^c)
}

// 按位置零操作符 &^
// z = x &^ y，表示如果 y 中的 bit 位为 1，则 z 对应 bit 位为 0，否则 z 对应 bit 位等于 x 中相应的 bit 位的值
// 或操作符 |  的效果与 &^ 完全相反
// z = x | y， 表示如果 y 中的 bit 位为 1，则 z 对应 bit 位为 1，否则 z 对应 bit 位等于 x 中相应的 bit 位的值
func TestZero(t *testing.T) {
	var x uint8 = 214
	var y uint8 = 92
	t.Logf("x: %08b\n", x)
	t.Logf("y: %08b\n", y)
	t.Logf("x | y: %08b", x|y)
	t.Logf("x &^ y: %08b", x&^y)
}
