package some

import (
	"strconv"
	"testing"
	"time"
)

// 整数转二进制
func convertI2B(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func TestConvertBinary(t *testing.T) {
	t.Log(convertI2B(12)) // 输出 1100
	t.Log(convertI2B(13)) // 输出 1101
}

func TestTimeout(t *testing.T) {
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			println("+1s")
		case <-tm:
			println("Timeout!!!")
			return
		}
	}
}

func TestArray(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	t.Log(a == nil)
}
