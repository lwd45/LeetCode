package mid

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := 0; i < len(edges); i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			return find(parent[x])
		}
		return x
	}

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		fromP, toP := find(from), find(to)
		if fromP != toP {
			parent[toP] = fromP
		} else {
			return edge
		}
	}

	return nil
}
