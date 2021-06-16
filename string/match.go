package string

//func isValid(s string) bool {
//	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
//	var stack []rune
//
//	for _, char := range s {
//		fmt.Println(reflect.TypeOf(char))
//		if char == '(' || char == '{' || char == '[' {
//			// 入栈
//			stack = append(stack, char)
//			// 循环中，stack不能为空
//		} else if len(stack) > 0 && brackets[char] == stack[len(stack) - 1] {
//			// 栈中有数据，且此元素与栈尾元素相同
//			stack = stack[:len(stack) - 1]
//		} else {
//			return false
//		}
//	}
//
//	// 循环结束，栈中还有数据则 false
//	return len(stack) == 0
//}

func isValid(strs string) bool {
	m := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	var stack []rune
	for _, str := range strs {
		if _, ok := m[str]; ok {
			stack = append(stack, str)
		} else if len(stack) > 0 && str == m[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}