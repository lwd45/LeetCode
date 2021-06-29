package easy

import "strconv"

func addStrings(num1 string, num2 string) string {
	len1, len2, ret := len(num1), len(num2), ""

	i, j, carry := len1-1, len2-1, 0
	for ; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		sum = sum % 10
		ret += strconv.Itoa(sum)
	}
	if carry == 1 {
		ret += "1"
	}

	bytes := []byte(ret)
	for i, index := 0, len(bytes)-1; i < index; i, index = i+1, index-1 {
		temp := bytes[i]
		bytes[i] = bytes[index]
		bytes[index] = temp
	}
	return string(bytes)
}

