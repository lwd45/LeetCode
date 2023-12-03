package mid

/*
[[4,1,3],[8,2,3],[1,3,6],[8,4,6],[4,4,8]]
12
*/
//func carPooling(trips [][]int, capacity int) bool {
//		sort.Slice(trips, func(i, j int) bool {
//			if trips[i][1] == trips[j][1] {
//				return trips[i][2] < trips[j][2]
//			}
//			return trips[i][1] < trips[j][1]
//		})
//
//		var save [][]int
//		for _, trip := range trips {
//			var saveTemp [][]int
//			for _, temp := range save {
//				if temp[0] <= trip[1] {
//					capacity += temp[1]
//				} else {
//					saveTemp = append(saveTemp, temp)
//				}
//			}
//			save = saveTemp
//			if capacity < trip[0] {
//				return false
//			}
//			capacity -= trip[0]
//			save = append(save, []int{trip[2], trip[0]})
//		}
//		return true
//}

func carPooling(trips [][]int, capacity int) bool {
		maxTo := 0
		for _, trip := range trips {
			maxTo = max(maxTo, trip[2])
		}

		diff := make([]int, maxTo+1)
		for _, trip := range trips {
			diff[trip[1]] += trip[0]
			diff[trip[2]] -= trip[0]
		}

		cnt := 0
		for i := 0; i <= maxTo; i++ {
			cnt += diff[i]
			if cnt > capacity {
				return false
			}
		}
		return true
}
