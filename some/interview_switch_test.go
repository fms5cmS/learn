package some

import "testing"

func TestSwitch(t *testing.T) {
	// value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// 无类型的常量 4 会被自动转换为这种常量的默认类型的值，这里是 int 类型的 4
	// switch 1 + 3 {
	// // case 中的类型为 int8，类型不同，所以无法通过编译！
	// case value1[0], value1[1]:  // case 可通过逗号分隔来列举相同的处理条件
	// 	t.Log("0 or 1")
	// case value1[2], value1[3]:
	// 	t.Log("2 or 3")
	// case value1[4], value1[5], value1[6]:
	// 	t.Log("4 or 5 or 6")
	// }
}

func TestSwitch_break(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 10:
			t.Log(10)
		case 2:
			t.Log(2)
			break
		}
		t.Logf("%d in for", i)
	}
	t.Log("out for")
}
