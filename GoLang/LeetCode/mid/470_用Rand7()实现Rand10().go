package mid

import (
	"math/rand"
)

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	var sum int
	for sum = (rand7()-1)*7 + rand7(); sum > 40; sum = (rand7()-1)*7 + rand7() {
	}
	return sum%10 + 1
}
