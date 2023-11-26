package mid
//
///*
//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//整数除法仅保留整数部分。
//
//输入：s = "3+2*2"
//输出：7
//
//输入：s = " 3/2 "
//输出：1
//
//输入：s = " 3+5 / 2 "
//输出：5
//
//1 <= s.length <= 3 * 105
//s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
//s 表示一个 有效表达式
//表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
//题目数据保证答案是一个 32-bit 整数
//*/
//func calculate(s string) int {
//	var stack []int
//	var singe uint8
//	singe = '+'
//
//	for i := 0; i < len(s); i++ {
//		if s[i] == ' ' {
//			continue
//		}
//
//		if s[i] <= '9' && s[i] >= '0' {
//			temp := 0
//			for i < len(s) && s[i] <= '9' && s[i] >= '0' {
//				temp = 10*temp + int(s[i]-'0')
//				i++
//			}
//			i--
//			if singe == '+' {
//				stack = append(stack, temp)
//			} else if singe == '-' {
//				stack = append(stack, -temp)
//			} else {
//				num := stack[len(stack)-1]
//				stack = stack[:len(stack)-1]
//				stack = append(stack, ans(num, singe, temp))
//			}
//		} else {
//			singe = s[i]
//		}
//	}
//
//	ret := 0
//	for i := 0; i < len(stack); i++ {
//		ret += stack[i]
//	}
//	return ret
//}
//
//func ans(num int, singe uint8, temp int) int {
//	if singe == '*' {
//		return num * temp
//	} else {
//		return num / temp
//	}
//}
