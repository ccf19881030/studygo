// 二分查找
package main

import "fmt"

// arr: 要查找的有序数组
// finddata: 要查找的元素
// 返回值：当finddata在arr数组中时，返回其在arr中的下标；如果不存在，则返回-1
func BinarySearch(arr []int, finddata int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if arr[mid] > finddata {
			high = mid - 1
		} else if arr[mid] < finddata {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	result := BinarySearch(arr, 12)
	if result == -1 {
		fmt.Println("没有找到12")
	} else {
		fmt.Println("12在 arr数组中的下标为：", result)
	}
}
