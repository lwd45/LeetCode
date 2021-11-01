package easy

func distributeCandies(candyType []int) int {
	length := len(candyType)
	set := make(map[int]struct{})
	for _, key := range candyType {
		if _, ok := set[key]; !ok {
			set[key] = struct{}{}
		}
	}
	if length/2 > len(set) {
		return len(set)
	}
	return length / 2
}
