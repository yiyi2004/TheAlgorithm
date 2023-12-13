package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//s := "3[a]2[bc]"
	res := decodeString("3[a]2[bc]")
	fmt.Println(res)
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
//func decodeString(s string) string {
//	if s == "" {
//		return ""
//	}
//	stack := make([]string, 0)
//	res := ""
//	for i := 0; i < len(s); i++ {
//		if s[i] >= '0' && s[i] <= '9' {
//			stack = append(stack, string(s[i]))
//		} else if s[i] == '[' {
//			stack = append(stack, string(s[i]))
//		} else if s[i] == ']' {
//			// 开始出栈
//			tmpStr := ""
//			stack = stack[:len(stack)-1]
//			for stack[len(stack)-1] != "[" {
//				tmpStr += stack[len(stack)-1]
//				stack = stack[:len(stack)-1]
//			}
//			stack = stack[:len(stack)-1]
//			n := ""
//			for stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
//				n += stack[len(stack)-1]
//				stack = stack[:len(stack)-1]
//			}
//			repeatString := getRepeatString(n, tmpStr)
//			stack = append(stack, repeatString)
//		} else {
//			// s[i] is character
//			stack = append(stack, string(s[i]))
//		}
//
//	}
//	return res
//}
//
//func getRepeatString(n string, str string) string {
//	n = reverseString(n)
//	atoi, _ := strconv.Atoi(n)
//	str = reverseString(str)
//	for i := 0; i < atoi; i++ {
//		str += str
//	}
//	return str
//}
//
//func reverseString(str string) string {
//	tmpStr := ""
//	for i := len(str) - 1; i >= 0; i-- {
//		tmpStr += string(str[i])
//	}
//	return tmpStr
//}

//func decodeString(s string) string {
//	stk := []string{}
//	ptr := 0
//	for ptr < len(s) {
//		cur := s[ptr]
//		if cur >= '0' && cur <= '9' {
//			digits := getDigits(s, &ptr)
//			stk = append(stk, digits)
//		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
//			stk = append(stk, string(cur))
//			ptr++
//		} else {
//			ptr++
//			sub := []string{}
//			for stk[len(stk)-1] != "[" {
//				sub = append(sub, stk[len(stk)-1])
//				stk = stk[:len(stk)-1]
//			}
//			// reverse
//			for i := 0; i < len(sub)/2; i++ {
//				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
//			}
//			// 取出 [
//			stk = stk[:len(stk)-1]
//			// 取出数字
//			repTime, _ := strconv.Atoi(stk[len(stk)-1])
//			stk = stk[:len(stk)-1]
//			// 重复字符串
//			t := strings.Repeat(getString(sub), repTime)
//			// 添加到 栈中
//			stk = append(stk, t)
//		}
//	}
//	return getString(stk)
//}
//
//func getDigits(s string, ptr *int) string {
//	ret := ""
//	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
//		ret += string(s[*ptr])
//	}
//	return ret
//}
//
//func getString(v []string) string {
//	ret := ""
//	for _, s := range v {
//		ret += s
//	}
//	return ret
//}

// 并不是什么困难的问题
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	var stk []string
	index := 0
	for index < len(s) {
		cur := s[index]
		if cur >= '0' && cur <= '9' {
			stk = append(stk, getDigits(s, &index))
		} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' || cur == '[' {
			stk = append(stk, string(cur))
			index++
		} else {
			index++
			tmp := []string{}
			for stk[len(stk)-1] != "[" {
				tmp = append(tmp, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			// 0 1 2 3 4 | 5
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
			}
			stk = stk[:len(stk)-1]
			n := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			str := ""
			for i := 0; i < len(tmp); i++ {
				str += tmp[i]
			}
			atoi, _ := strconv.Atoi(n)
			stk = append(stk, strings.Repeat(str, atoi))
		}
	}
	return getString(stk)
}

func getDigits(s string, index *int) string {
	// 如果是数字，就一直循环
	res := ""
	for s[*index] >= '0' && s[*index] <= '9' {
		res += string(s[*index])
		*index++
	}
	return res
}

func getString(stk []string) string {
	res := ""
	for _, v := range stk {
		res += v
	}
	return res
}
