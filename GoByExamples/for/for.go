package main

import "fmt"

func main() {
  
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // 经典的初始化/条件/后续形式for循环
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("loop")
    break
  }

}
