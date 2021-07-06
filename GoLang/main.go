package main

import "fmt"

func main() {
	//fmt.Println(mid.FrequencySort("asdas"))
	buildTree(make([]int, 0))
}

func buildTree(preorder []int) {
	fmt.Printf("%T\n", preorder)
}
