package some

import "testing"

/* 易错的 */

func TestNilCover(t *testing.T) {
	nil := 123
	t.Log(nil)
	// var _ map[int]string = nil
}

func TestNil(t *testing.T) {
	var m map[int]bool
	_ = m[123]
	var p *[5]string
	for range p {
		_ = len(*p)
	}
	var s []int
	_ = s[:]
	// s[0] 的 s == nil ，报错
	s, s[0] = []int{1, 2}, 0
}

func TestOverflow(t *testing.T) {
	var a int8 = -128
	t.Log(a / -1) // -128，溢出了！
}

func TestHex(t *testing.T) {
	const (
		century = 100
		decade  = 010
		year    = 001
	)
	// decade 和 year 是八进制，在计算时会转为十进制
	t.Log(century + 2*decade + 2*year)  // 118
}

func TestSelect(t *testing.T) {
	var o = t.Log
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c:
			o(2)
			c = nil
		case c <- 1:
			o(3)
		}
	}
}
