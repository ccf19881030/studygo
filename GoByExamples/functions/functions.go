package main

import "fmt"

func add(a int, b int) int {
  return a + b
}

func sub(a int, b int) int {
  return a - b
}

func plusThreeNumbers(a, b, c int) int {
   return a + b + c
}

func main() {
  res := add(1, 2)
  fmt.Println("1+2 = ", res)

  res = sub(4, 5)
  fmt.Println("4-5 = ", res)

  res = plusThreeNumbers(3, 5, 10)
  fmt.Println("3+5+10 = ", res)
}
