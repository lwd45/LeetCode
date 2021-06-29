package easy

func isValid(s string) bool {
	arr := make([]int32, 0)
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			arr = append(arr, c)
		case ')':
			if len(arr) == 0 || arr[len(arr)-1] != '(' {
				return false
			} else {
				arr = arr[:len(arr)-1]
			}
		case ']':
			if len(arr) == 0 || arr[len(arr)-1] != '[' {
				return false
			} else {
				arr = arr[:len(arr)-1]
			}
		case '}':
			if len(arr) == 0 || arr[len(arr)-1] != '{' {
				return false
			} else {
				arr = arr[:len(arr)-1]
			}
		}
	}
	return len(arr) == 0
}
