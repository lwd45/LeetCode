package main

import "fmt"

func main() {
	s := "01010"
	for _, c := range s {
		fmt.Printf("%v", c)
	}
}
