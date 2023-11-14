package main

import (
	"strconv"
)

func main() {
	//s := "3[a]2[bc]"
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

// 可以，因为没有指定下标，所以完全错了。
func decodeString(s string) string {
	if s == "" {
		return ""
	}
	stack := make([]string, 0)
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			stack = append(stack, string(s[i]))
		} else if s[i] == '[' {
			stack = append(stack, string(s[i]))
		} else if s[i] == ']' {
			// 开始出栈
			tmpStr := ""
			stack = stack[:len(stack)-1]
			for stack[len(stack)-1] != "[" {
				tmpStr += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			n := ""
			for stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				n += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			repeatString := getRepeatString(n, tmpStr)
			stack = append(stack, repeatString)
		} else {
			// s[i] is character
			stack = append(stack, string(s[i]))
		}

	}
	return res
}

func getRepeatString(n string, str string) string {
	n = reverseString(n)
	atoi, _ := strconv.Atoi(n)
	str = reverseString(str)
	for i := 0; i < atoi; i++ {
		str += str
	}
	return str
}

func reverseString(str string) string {
	tmpStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		tmpStr += string(str[i])
	}
	return tmpStr
}
