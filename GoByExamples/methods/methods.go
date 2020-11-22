// Go by Example 中文： 方法
// Go 支持在结构体类型中定义方法
package main

import "fmt"

// 定义一个矩形结构体
type rectangle struct {
    width int
    height int
}

// 可以为值类型或者指针类型的接收器定义方法。
// 矩形面积
// 这里的area方法有一个接收器类型rectangle
func (r *rectangle) area() int {
  return r.width * r.height
}

// 这里是一个值接收器的例子。
// 矩形周长
func (r rectangle) perim() int {
  return 2 * r.width + 2 * r.height
}

func main() {
  r := rectangle{width: 10, height: 5}

  fmt.Println("area: ", r.area())
  fmt.Println("perim: ", r.perim())
	
  // Go 自动处理方法调用时的值和指针之间的转化。你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或中让方法能够改变接受的数据。
  rp := &r
  fmt.Println("area: ", rp.area())
  fmt.Println("perim: ", rp.perim())
}

