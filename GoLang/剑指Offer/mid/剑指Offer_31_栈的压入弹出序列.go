package mid

//func validateStackSequences(pushed []int, popped []int) bool {
//	var stack []int
//	pushIndex, popIndex := 0, 0
//
//	for popIndex < len(popped) {
//		if len(stack) <= 0 || stack[len(stack)-1] != popped[popIndex] {
//			if pushIndex >= len(pushed) {
//				return false
//			}
//			stack = append(stack, pushed[pushIndex])
//			pushIndex++
//		} else {
//			stack = stack[:len(stack)-1]
//			popIndex++
//		}
//	}
//	return true
//}

func validateStackSequences(pushed []int, popped []int) bool {
	pushIndex, popIndex := 0, 0
	for _, v := range pushed {
		pushed[pushIndex] = v
		for pushIndex >= 0 && popped[popIndex] == pushed[pushIndex] {
			pushIndex--
			popIndex++
		}
		pushIndex++
	}
	return popIndex == len(popped)
}
