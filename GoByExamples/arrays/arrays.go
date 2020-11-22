// Go by Example 中文： 数组
// https://books.studygolang.com/gobyexample/arrays/
package main

import "fmt"

func main() {
   var a [5]int
   fmt.Println("emp:", a)

   a[4] = 100
   fmt.Println("set:", a)
   fmt.Println("get:", a[4])

   // 使用内置函数len返回数组的长度
   fmt.Println("len:", len(a))

   // 使用这个语法在一行内初始化一个数组
   b := [5]int{1, 2, 3, 4, 5}
   fmt.Println("dcl:", b)

   // 数组的存储类型是单一的，但是你可以组合这些数据结构来构造多维的数据结构
   var twoD [2][3]int
   for i := 0; i < 2; i++ {
	for j := 0; j < 3; j++ {
   	    twoD[i][j] = i + j
	}
   }
   fmt.Println("2d: ", twoD)
}
