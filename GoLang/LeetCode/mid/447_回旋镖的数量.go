package mid

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for _, point1 := range points {
		dist := make(map[int]int)
		for _, point2 := range points {
			dis := (point1[0]-point2[0])*(point1[0]-point2[0]) + (point1[1]-point2[1])*(point1[1]-point2[1])
			dist[dis]++
		}

		for _, v := range dist {
			res += v*(v-1)
		}
	}
	return res
}
