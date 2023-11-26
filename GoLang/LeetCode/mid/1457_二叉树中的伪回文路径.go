package mid

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pseudoPalindromicPaths(root *TreeNode) int {
	ans, oldCnt := 0, 0
	cntMap := make(map[int]int, 10)

	dfs_1457(root, &ans, &oldCnt, cntMap)
	return ans
}

func dfs_1457(root *TreeNode, ans, cnt *int, cntMap map[int]int) {
	if root == nil {
		return
	}

	cntMap[root.Val]++
	if cntMap[root.Val]%2 == 0 {
		*cnt--
	} else {
		*cnt++
	}

	dfs_1457(root.Left, ans, cnt, cntMap)
	dfs_1457(root.Right, ans, cnt, cntMap)

	if root.Left == nil && root.Right == nil && *cnt <= 1 {
		*ans++
	}

	cntMap[root.Val]--
	if cntMap[root.Val]%2 == 0 {
		*cnt--
	} else {
		*cnt++
	}
}
