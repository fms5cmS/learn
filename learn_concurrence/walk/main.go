package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "E:\\mod\\src\\learn\\learn_concurrence"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("the error is %s, current path is %s\n", err, path)
			return err
		}
		fmt.Println(path)
		if info.IsDir() {
			fmt.Printf("the dir is %s \n", info.Name())
		} else {
			fmt.Printf("the file size is %d \n", info.Size())
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
