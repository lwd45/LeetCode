package mid

func garbageCollection(garbage []string, travel []int) int {
	g, m, p := 0, 0, 0
	gt, mt, pt := 0, 0, 0
	for i, s := range garbage {
		for _, c := range s {
			if c == 'G' {
				g += 1 + gt
				gt = 0
			} else if c == 'M' {
				m += 1 + mt
				mt = 0
			} else if c == 'P' {
				p += 1 + pt
				pt = 0
			}
		}

		if i != len(garbage)-1 {
			gt += travel[i]
			mt += travel[i]
			pt += travel[i]
		}
	}
	return g + m + p
}
