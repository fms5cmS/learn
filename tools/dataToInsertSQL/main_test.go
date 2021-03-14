package main

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	fields, values := parseFile("player.txt")
	fmt.Println("fields: ", fields)
	fmt.Println("values: ")
	for _, value := range values {
		fmt.Println(value)
	}
}

func TestGenerateStr(t *testing.T) {
	strs := []string{"1001", "1001", "aaa", "23.90"}
	t.Log(generateStr(strs, "value"))
}

func TestAll(t *testing.T) {
	fields, values := parseFile("player.txt", 3)
	generateSQL("player", fields, values)
}
