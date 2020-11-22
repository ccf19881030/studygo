package main

import "fmt"
import "io/ioutil"

func main() {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("read file error")
		return
	}
	fmt.Println(string(b))
}
