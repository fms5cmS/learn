package some

import (
	"fmt"
	"strings"
	"testing"
)

// 编译无法通过，nil 无法赋值给 string 类型
// func GetValue(m map[int]string,id int) (string,bool) {
// 	if _,ok:=m[id];ok{
// 		return "exist",true
// 	}
// 	return nil,false
// }

func TestStrConn(t *testing.T) {
	t.Log("test the connection of strings")
	// 方式一
	str1 := "abc" + "123"
	t.Log(str1)
	// 方式二
	str2 := fmt.Sprintf("abc%d", 123)
	t.Log(str2)
	// 方式三
	strs := []string{"abc", "1", "23"}
	// 第二个参数是连接每个字符串的连接符
	str3 := strings.Join(strs, "")
	t.Log(str3)
	// 方式四
	str4 := strings.Builder{}
	str4.WriteString("abc")
	str4.WriteString("123")
	t.Log(str4.String())
}

func TestStrChange(t *testing.T) {
	str := "proto"
	t.Log(str[0])                                      // 这里输出的是 104 对应的字符为 'h'
	t.Logf("%T,%v,%s", str[0], str[0], string(str[0])) // uint8,104,h
	// str[0] = 'x'  字符串是只读的，编译报错
	t.Log(str)
}

func TestUTF(t *testing.T) {
	str := "fS 组合"
	for i, v := range str {
		// %q：对 string 或 []byte 类型的值输出时：带双引号、带反引号
		// %x：输出小写十六进制数，% x 对两个字节的十六进制数含空格输出
		t.Logf("%d: %q [% x]", i, v, []byte(string(v)))
	}
}

func TestUtf(t *testing.T) {
	str := "fS 组合"
	t.Logf("The string: %q\n", str)
	// 打印单个 Unicode 字符
	// %q：对 string 或 []byte 类型的值输出时带双/单引号
	t.Logf(" => runes(char): %q\n", []rune(str))
	// 每个 rune 类型的值底层都是使用 UTF-8 编码值表示的
	// %x：输出小写十六进制数，% x 对两个字节的十六进制数含空格输出
	t.Logf(" => runes(hex): %x\n", []rune(str))
	// 中文字符对应的 UFT-8 编码都需要使用三个字节表示，将每个字符的 UTF-8 编码值拆成对应的字节序列
	t.Logf(" => bytes(hex): [% x]\n", []byte(str))
}

func TestStr(t *testing.T) {
	str := "123"
	([]byte)(str)[1] = 'b'
	t.Log(str)
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		go func(kk *string, vv *int) {
			fmt.Printf(" %v , %v \n", *kk, *vv)
		}(&k, &v) // d
	}
}
