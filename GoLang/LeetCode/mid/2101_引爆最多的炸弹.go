package mid

import "container/list"

func maximumDetonation(bombs [][]int) int {
	edges := make(map[int][]int)

	for i := 0; i < len(bombs); i++ {
		for j := 0; j < len(bombs); j++ {
			dx := bombs[i][0] - bombs[j][0]
			dy := bombs[i][1] - bombs[j][1]
			r := bombs[i][2]

			if j != i && dx*dx+dy*dy <= r*r {
				edges[i] = append(edges[i], j)
			}
		}
	}

	res := 0
	for i := 0; i < len(bombs); i++ {
		visited := make([]bool, len(bombs))
		cnt := 1

		visited[i] = true
		q := list.New()
		q.PushBack(i)
		for q.Len() > 0 {
			remove := q.Remove(q.Front()).(int)
			for _, v := range edges[remove] {
				if !visited[v] {
					q.PushBack(v)
					visited[v] = true
					cnt++
				}
			}
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}
