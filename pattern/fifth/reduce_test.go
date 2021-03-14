package fifth

import (
	"fmt"
	"testing"
)

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func TestReduce(t *testing.T) {
	var list = []string{"Hao", "Chen", "MegaEase"}
	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x)  // 15
}
