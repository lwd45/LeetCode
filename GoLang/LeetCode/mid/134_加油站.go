package mid

/*
class Solution {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int length = gas.length;
        int i = 0;
        while (i < length) {
            int forward = 0;
            int gasCount = 0;
            int costCount = 0;
            while (forward < length) {
                gasCount += gas[(i + forward) % length];
                costCount += cost[(i + forward) % length];
                if (gasCount < costCount)
                    break;
                forward += 1;
            }
            if (forward == length)
                return i;
            else
                i = i + forward + 1;
        }
        return -1;
    }
}
*/

func canCompleteCircuit(gas []int, cost []int) int {
	start, length := 0, len(gas)
	for start < length {
		offset, gasCount, costCount := 0, 0, 0
		for offset < length {
			gasCount += gas[(start+offset)%length]
			costCount += cost[(start+offset)%length]
			if gasCount < costCount {
				break
			}
			offset++
		}

		if offset == length {
			return start
		} else {
			start = start + offset + 1
		}
	}
	return -1
}

//func canCompleteCircuit(gas []int, cost []int) int {
//	left, minLeft, insertIndex := 0, math.MaxInt64, -1
//	for i := 0; i < len(gas); i++ {
//		left += gas[i] - cost[i]
//		if left < minLeft {
//			insertIndex = i
//			minLeft = left
//		}
//	}
//	if left < 0 {
//		return -1
//	}
//	return (insertIndex + 1) % len(gas)
//}
