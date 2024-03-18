package mid

import "fmt"

//func getHint(secret string, guess string) string {
//	A, B := 0, 0
//	mapA := map[uint8]int{}
//	mapB := map[uint8]int{}
//	for i, _ := range secret {
//		if secret[i] == guess[i] {
//			A++
//		} else {
//			mapA[secret[i]] = mapA[secret[i]] + 1
//			mapB[guess[i]] = mapB[guess[i]] + 1
//		}
//	}
//
//	for key, _ := range mapA {
//		if value, ok := mapB[key]; ok {
//			B += min(mapA[key], value)
//		}
//	}
//	return fmt.Sprintf("%dA%dB", A, B)
//}

func getHint(secret string, guess string) string {
	aCnt, bCnt := make(map[uint8]int, len(secret)), make(map[uint8]int, len(guess))

	A, B := 0, 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			A++
		} else {
			aCnt[secret[i]] ++
			bCnt[guess[i]] ++
		}
	}

	for key, cntA := range aCnt {
		if cntB, ok := bCnt[key]; ok {
			B += min(cntA, cntB)
		}
	}
	return fmt.Sprintf("%vA%vB", A, B)
}
