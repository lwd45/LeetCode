package mid

//func countTriplets(arr []int) int {
//	length, S := len(arr), make([]int, len(arr)+1)
//	for i := 0; i < length; i++ {
//		S[i+1] = S[i] ^ arr[i]
//	}
//
//	ans := 0
//	for i := 0; i < length; i++ {
//		for j := i + 1; j < length; j++ {
//			for k := j; k < length; k++ {
//				if S[i] == S[k+1] {
//					ans++
//				}
//			}
//		}
//	}
//	return ans
//}

//func countTriplets(arr []int) int {
//	length, S := len(arr), make([]int, len(arr)+1)
//	for i := 0; i < length; i++ {
//		S[i+1] = S[i] ^ arr[i]
//	}
//
//	ans := 0
//	for i := 0; i < length; i++ {
//		for k := i + 1; k < length; k++ {
//			if S[i] == S[k+1] {
//				ans += k - i
//			}
//		}
//	}
//	return ans
//}

// S[i] = a[0] ^ a[1] ^ ... a[i-1]
// S[j] = a[0] ^ a[1] ^ ... a[j-1]
// S[k] = a[0] ^ a[1] ^ ... a[k-1]

// a[i] ^ a[i+1] ^ ... a[j-1] = S[i] ^ S[j] = a
// a[j] ^ a[j+1] ^ ... a[k] = S[j] ^ S[k+1] = b
// a == b    --->    S[i] == S[k+1]
func countTriplets(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		for k := i + 1; k < len(arr); k++ {
			sum ^= arr[k]
			if sum == 0 {
				count += k - i
			}
		}
	}
	return count
}
