package main

func main() {

}

//type stack []byte
//
//func (s *stack) peek() uint8 {
//	return (*s)[len(*s)-1]
//}
//
//func (s *stack) push(c uint8) {
//	*s = append(*s, c)
//}
//
//func (s *stack) pop() uint8 {
//	v := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return v
//}
//
//func (s *stack) isEmpty() bool {
//	return len(*s) == 0
//}
//
//func isLeft(c uint8) bool {
//	if c == '(' || c == '{' || c == '[' {
//		return true
//	}
//	return false
//}
//
//func isTwo(c1, c2 uint8) bool {
//	if c1 == '(' && c2 == ')' {
//		return true
//	}
//	if c1 == '{' && c2 == '}' {
//		return true
//	}
//	if c1 == '[' && c2 == ']' {
//		return true
//	}
//	return false
//}
//
//func isValid(s string) bool {
//	if len(s)%2 != 0 {
//		return false
//	}
//	var st stack
//	for i := 0; i < len(s); i++ {
//		if isLeft(s[i]) {
//			st.push(s[i])
//		} else {
//			if st.isEmpty() {
//				return false
//			} else {
//				c := st.peek()
//				if isTwo(c, s[i]) {
//					st.pop()
//				} else {
//					return false
//				}
//			}
//		}
//	}
//	if st.isEmpty() {
//		return true
//	}
//	return false
//}

// 其实是一样的
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
