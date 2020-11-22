// Go by Example 中文版: 指针
// https://gobyexample-cn.github.io/pointers
// Go 支持指针，允许在程序中通过引用传递来传递值和数据结构

package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

// 通过值传递是无法交换两个整数的
func swapByValue(a int, b int) {
    var temp int
    temp = a
    a = b
    b = temp
}

// 通过指针可以交换两个整数
func swapByPointer(pa *int, pb *int) {
    var pTemp int
    pTemp = *pa
    *pa = *pb
    *pb = pTemp
}

func main() {
   i := 1
   fmt.Println("initial:", i)

   zeroval(i)
   fmt.Println("zeroval:", i)

   zeroptr(&i)
   fmt.Println("zeroptr:", i)

   // 指针也是可以被打印的
   fmt.Println("pointer address:", &i)

   var c, d = 13, 45
   fmt.Println("c=",c,"d=",d)
   swapByValue(c, d)
   fmt.Println("after swapByValue is called, c=", c, "d=", d)
   swapByPointer(&c, &d)
   fmt.Println("after swapByPointer is called, c=", c, "d=", d)
}
