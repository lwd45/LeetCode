package mid

import "strings"

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}

	builder := strings.Builder{}
	for n != 0 {
		r := n % -2
		if r == 0 {
			builder.WriteString("0")
		} else {
			n -= 1
			builder.WriteString("1")
		}
		n /= -2
	}
	return builder.String()
}

func reverseString(s string) string {
	// 将字符串转换为字节数组
	bytes := []byte(s)
	// 获取字节数组的长度
	length := len(bytes)
	// 使用双指针进行字符交换
	left, right := 0, length-1
	for left < right {
		// 交换左右指针对应的字符
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
	// 返回翻转后的字符串
	return string(bytes)
}
