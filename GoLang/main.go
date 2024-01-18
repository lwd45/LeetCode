package main

func main() {
	addMinimum("aaaaba")
}

func addMinimum(word string) int {
	n := len(word)

	res := 1
	for i := 1; i < n; i++ {
		k := int(word[i]) - int(word[i-1])
		if k <= 0 {
			res++
		}
	}

	return 3*res - n
}
