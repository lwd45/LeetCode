package mid

func maxProduct(words []string) int {
	mask := make([]int, len(words))
	for idx, word := range words {
		t := 0
		for _, c := range word {
			v := c - 'a'
			t |= 1 << v // 1左移v位  & | ^
		}
		mask[idx] = t
	}
	ans := 0
	for i := 0; i < len(mask); i++ {
		for j := 0; j < i; j++ {
			if mask[i]&mask[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}
