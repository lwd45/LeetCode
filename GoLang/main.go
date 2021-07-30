package main

import "fmt"

func main() {
	str := "luoweidong"
	fmt.Println(str)
	fmt.Println("-----------------")

	for _, v := range str {
		fmt.Println(v)
	}
	fmt.Println("-----------------")

	cs := []byte(str)
	for _, c := range cs {
		fmt.Println(byte(c))
	}
	fmt.Println("-----------------")
	fmt.Println(cs)
}
