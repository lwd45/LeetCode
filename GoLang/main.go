package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	//a := []int{5, 4, 3, 2, 1}
	c := sort.Search(len(a), func(i int) bool { return a[i] < 3 })
	fmt.Println(c)
}
