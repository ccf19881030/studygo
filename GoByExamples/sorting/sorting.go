package main

import (
    "fmt"
    "sort"
)

func main() {
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    // 一个int排序的例子。
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints: ", ints)

    // 我们也可以使用sort来检查一个切片是否为有序的。
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
