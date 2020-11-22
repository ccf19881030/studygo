package main
// Go by Example 中文版: 数字解析
// https://gobyexample-cn.github.io/number-parsing
// 从字符串中解析数字在很多程序中是一个基础常见的任务， 而在 Go 中，是这样处理的。
// 内建的strconv包提供了数字解析能力。
import (
	"fmt"
	"strconv"
)

func main() {
	// 使用ParseFloat，这里的64表示解析的数的位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 在使用ParseFloat解析整型数时，例子中的参数0表示自动腿短字符串所表示的数字的数制。
	// 64表示返回的整型数是以64位存储的。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt 会自动识别出字符串是十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint 也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi 是一个基础的10进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

