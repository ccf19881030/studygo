package main

// Go by Example 中文版: SHA1 哈希
// https://gobyexample-cn.github.io/sha1-hashes

//SHA1 散列（hash） 经常用于生成二进制文件或者文本块的短标识。
//例如，git 版本控制系统 大量的使用了 SHA1 来标识受版本控制的文件和目录。
//这是 Go 中如何进行 SHA1 散列计算的例子。

// Go在多个crypto/*包中实现了一系列散列函数
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// 产生一个散列值的方式是 sha1.New()
	// sha1.Wrtie(bytes)，然后 sha1.Sum([]bytes{})。这里我们从一个新的散列开始。
	h := sha1.New()

	// 写入要处理的字节。如果是一个字符串，需要使用[]byte(s)将其强制转换成字节数组。
	h.Write([]byte(s))

	// Sum得到最终的散列值的字符切片。Sum接收一个参数，可以用来给现有的字符切片追加额外的字符切片：
	// 但是一般都不需要这样做。
	bs := h.Sum(nil)

	// SHA1值经常以16进制输出，例如在 git commit中。我们这里也使用%x来将散列结果格式化为16进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
