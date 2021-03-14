package fifth

import (
	"fmt"
	"strings"
	"testing"
)

func MapStrToInt(arr []string, fn func(s string) int) []int {
	newArray := make([]int, 0, len(arr))
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToStr(arr []string, fn func(s string) string) []string {
	newArray := make([]string, 0, len(arr))
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func TestMap(t *testing.T) {
	var list = []string{"Hao", "Chen", "MegaEase"}
	
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
	// ["HAO", "CHEN", "MEGAEASE"]
}
