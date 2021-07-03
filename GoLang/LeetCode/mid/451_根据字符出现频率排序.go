package mid

import (
	"bytes"
	"sort"
)

/*
给定一个字符串，请将字符串里的字符按照出现的频率降序排列。

示例 1:
输入:
"tree"
输出:
"eert"

示例 2:
输入:
"cccaaa"
输出:
"cccaaa"

示例 3:
输入:
"Aabb"
输出:
"bbAa"


func frequencySort(s string) string {
    cnt := map[byte]int{}
    for i := range s {
        cnt[s[i]]++
    }

    type pair struct {
        ch  byte
        cnt int
    }
    pairs := make([]pair, 0, len(cnt))
    for k, v := range cnt {
        pairs = append(pairs, pair{k, v})
    }
    sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })

    ans := make([]byte, 0, len(s))
    for _, p := range pairs {
        ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
    }
    return string(ans)
}
*/
func frequencySort(s string) string {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}

	type pair struct {
		ch  byte
		cnt int
	}

	pairs := make([]pair, 0, len(m))
	for k, v := range m {
		p := pair{ch: k, cnt: v}
		pairs = append(pairs, p)
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })
	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}
