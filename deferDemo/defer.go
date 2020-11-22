package main

// Go by Example 中文版：Defer
// Defer用于确保程序在执行完成后，会调用某个函数，一般是执行清理工作。Defer的用途跟其他语言的ensure或finally类似。

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("D://defer.txt")
	defer closeFile(f)
	writeFile(f)
}

// 创建文件
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// 向文件中写入数据
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// 关闭文件
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}