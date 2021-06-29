package easy

/*
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

I可以放在V(5) 和X(10) 的左边，来表示 4 和 9。
X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
C可以放在D(500) 和M(1000) 的左边，来表示400 和900。
给定一个罗马数字，将其转换成整数。输入确保在 1到 3999 的范围内。
*/
func romanToInt(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+1 < len(s) && s[i+1] == 'V' {
				ans += 4
				i++
			} else if i+1 < len(s) && s[i+1] == 'X' {
				ans += 9
				i++
			} else {
				ans += 1
			}
		case 'V':
			ans += 5
		case 'X':
			if i+1 < len(s) && s[i+1] == 'L' {
				ans += 40
				i++
			} else if i+1 < len(s) && s[i+1] == 'C' {
				ans += 90
				i++
			} else {
				ans += 10
			}
		case 'L':
			ans += 50
		case 'C':
			if i+1 < len(s) && s[i+1] == 'D' {
				ans += 400
				i++
			} else if i+1 < len(s) && s[i+1] == 'M' {
				ans += 900
				i++
			} else {
				ans += 100
			}
		case 'D':
			ans += 500
		case 'M':
			ans += 1000
		}
	}
	return ans
}
