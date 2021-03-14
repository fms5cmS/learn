package first

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	
	dir1 := path[:sepIndex]
	// dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	
	fmt.Println("dir1 =>", string(dir1), len(dir1), cap(dir1)) // dir1 => AAAA 4 14
	fmt.Println("dir2 =>", string(dir2), len(dir2), cap(dir2)) // dir2 => BBBBBBBBB 9 9
	
	dir1 = append(dir1, "suffix"...)
	
	fmt.Println("dir1 =>", string(dir1), len(dir1), cap(dir1)) // dir1 => AAAAsuffix 10 14
	fmt.Println("dir2 =>", string(dir2), len(dir2), cap(dir2)) // dir2 => uffixBBBB 9 9
}

func TestSubSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[:6]
	t.Logf("s = %v, len = %d, cap = %d", s, len(s), cap(s))
	t.Log(s[7])
}
