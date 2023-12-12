package main

import (
	"fmt"
)

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}

//func generateParenthesis(n int) (res []string) {
//	if n == 0 {
//		return nil
//	}
//	tmp := make([]byte, 0, n*2)
//	var dfs func(index int)
//	dfs = func(index int) {
//		if index == 2*n {
//			if check(tmp) {
//				res = append(res, string(tmp))
//			}
//			return
//		}
//		tmp = append(tmp, '(')
//		dfs(index + 1)
//		tmp := tmp[:len(tmp)-1]
//		tmp = append(tmp, ')')
//		dfs(index + 1)
//	}
//	dfs(0)
//	return res
//}
//
//func check(str []byte) bool {
//	var stack string
//	for i := 0; i < len(str); i++ {
//		if str[i] == '(' {
//			stack = stack + string('(')
//		} else if str[i] == ')' {
//			if stack == "" || stack[len(stack)-1] != '(' {
//				return false
//			} else {
//				stack = stack[:len(stack)-1]
//			}
//		}
//	}
//	return true
//}

func generateParenthesis(n int) (res []string) {
	// 前面是 index， 后面是判断条件
	if n == 0 {
		return nil
	}
	tmp := make([]byte, n*2)
	var dfs func(index int, open int)
	dfs = func(index int, open int) {
		if index == n*2 {
			res = append(res, string(tmp))
			return
		}
		if open < n {
			tmp[index] = '('
			dfs(index+1, open+1)
		}
		if index-open < open {
			tmp[index] = ')'
			dfs(index+1, open)
		}
	}
	dfs(0, 0)
	return res
}
