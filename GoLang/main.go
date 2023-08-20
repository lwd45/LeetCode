package main

func main() {
	s := make([]int, 5)
	for i := 0; i < 8; i++ {
		s = append(s, i)
	}
	println(s)
}
