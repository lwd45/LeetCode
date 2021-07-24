package mid

func permutation(s string) []string {
	var ans []string
	var oneStr []byte
	var visited = make([]bool, len(s))
	var contains = make(map[string]struct{})

	dfs(s, 0, &ans, &oneStr, visited, contains)
	return ans
}

func dfs(s string, idx int, ans *[]string, oneStr *[]byte, visited []bool, contains map[string]struct{}) {
	if idx == len(s) {
		if _, ok := contains[string(*oneStr)]; !ok {
			*ans = append(*ans, string(*oneStr))
			contains[string(*oneStr)] = struct{}{}
		}
		return
	}

	for i := 0; i < len(s); i++ {
		if !visited[i] {
			visited[i] = true
			*oneStr = append(*oneStr, s[i])
			idx++
			dfs(s, idx, ans, oneStr, visited, contains)
			idx--
			*oneStr = (*oneStr)[:len(*oneStr)-1]
			visited[i] = false
		}
	}
}
