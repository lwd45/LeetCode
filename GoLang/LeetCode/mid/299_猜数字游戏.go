package mid

import "fmt"

func getHint(secret string, guess string) string {
	A, B := 0, 0
	mapA := map[uint8]int{}
	mapB := map[uint8]int{}
	for i, _ := range secret {
		if secret[i] == guess[i] {
			A++
		} else {
			mapA[secret[i]] = mapA[secret[i]] + 1
			mapB[guess[i]] = mapB[guess[i]] + 1
		}
	}

	for key, _ := range mapA {
		if value, ok := mapB[key]; ok {
			B += min(mapA[key], value)
		}
	}
	return fmt.Sprintf("%dA%dB", A, B)
}
