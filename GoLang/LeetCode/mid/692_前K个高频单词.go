package mid

import "sort"

/*
给一非空的单词列表，返回前k个出现次数最多的单词。
返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。

示例 1：
输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]
解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。注意，按字母顺序 "i" 在 "love" 之前。

示例 2：
输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
输出: ["the", "is", "sunny", "day"]
解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，出现次数依次为 4, 3, 2 和 1 次。

扩展练习：
尝试以O(n log k) 时间复杂度和O(n) 空间复杂度解决。
*/
func topKFrequent(words []string, k int) []string {
	wordCount := make(map[string]int, len(words))
	for _, word := range words {
		wordCount[word] = wordCount[word] + 1
	}

	uniqueWord := make([]string, 0, len(wordCount))
	for key, _ := range wordCount {
		uniqueWord = append(uniqueWord, key)
	}

	sort.Slice(uniqueWord, func(i, j int) bool {
		s1, s2 := uniqueWord[i], uniqueWord[j]
		return (wordCount[s1] > wordCount[s2]) || (wordCount[s1] == wordCount[s2] && s1 < s2)
	})
	return uniqueWord[:k]
}
