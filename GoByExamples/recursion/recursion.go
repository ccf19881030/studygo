// Go by Example 中文版: 递归
// https://gobyexample-cn.github.io/recursion
// Go 支持递归。这里是一个经典的阶乘示例。
package main

import "fmt"

// 阶乘函数
// fact 函数在到达 fact(0)前一直调用自身
func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}

// 斐波拉契数列
func fabniooc(n int) int {
  if n == 0 || n == 1 {
    return 1
  }  else {
    return fabniooc(n-1) + fabniooc(n-2)
  }
}

func main() {
  fmt.Println(fact(7))

  fmt.Println(fabniooc(4))
}
