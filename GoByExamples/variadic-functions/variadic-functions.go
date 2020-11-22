package main

import "fmt"

// Go by Example 中文版: 变参函数
// https://gobyexample-cn.github.io/variadic-functions
// 可变参数函数。 在调用时可以传递任意数量的参数。 例如，fmt.Println 就是一个常见的变参函数。

// sum 这个函数接受任意数量的int作为参数
func sum(nums ...int) {
  fmt.Println(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

// 变参函数使用常规的调用方式，传入独立的参数。
func main() {
  sum(1, 2)
  sum(1, 2, 3)

  // 如果你有一个含有多个值的 slice，想把它们作为参数使用，你需要这样调用 func(slice...)
  nums := []int{1, 2, 3, 4}
  sum(nums...)
}
