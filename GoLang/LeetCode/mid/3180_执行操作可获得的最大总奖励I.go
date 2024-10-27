package mid

import "sort"

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	m := rewardValues[len(rewardValues)-1]

	dp := make([]int, 2*m)
	dp[0] = 1

	for _, value := range rewardValues {
		for i := value*2 - 1; i >= value; i-- {
			if dp[i-value] == 1 {
				dp[i] = 1
			}
		}
	}

	for i := len(dp) - 1; i >= 0; i-- {
		if dp[i] == 1 {
			return i
		}
	}

	return 0
}
