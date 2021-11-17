package mid

import "sort"

/*
[[4,1,3],[8,2,3],[1,3,6],[8,4,6],[4,4,8]]
12
*/
func carPooling(trips [][]int, capacity int) bool {
		sort.Slice(trips, func(i, j int) bool {
			if trips[i][1] == trips[j][1] {
				return trips[i][2] < trips[j][2]
			}
			return trips[i][1] < trips[j][1]
		})

		var save [][]int
		for _, trip := range trips {
			var saveTemp [][]int
			for _, temp := range save {
				if temp[0] <= trip[1] {
					capacity += temp[1]
				} else {
					saveTemp = append(saveTemp, temp)
				}
			}
			save = saveTemp
			if capacity < trip[0] {
				return false
			}
			capacity -= trip[0]
			save = append(save, []int{trip[2], trip[0]})
		}
		return true
}
