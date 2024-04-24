package main

import "fmt"

func main() {
	s := ")()())"
	res := longestValidParentheses(s)
	fmt.Println(res)
}

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	// dp[i] 表示 [0:i] 有多少对
	dp := make([]int, len(s))
	stack := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, "(")
		} else {
			if len(stack) == 0 {
				continue
			}
			// 只有栈顶是 "(" 的时候才会构成一对
			if stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
				// pairs
				pairs := dp[i-1] + 1
				preIndex := i - 2*pairs
				if preIndex > 0 {
					dp[i] = pairs + dp[preIndex]
				} else {
					dp[i] = pairs
				}
			}
		}
	}
	return max(dp) * 2
}

func longestValidParentheses2(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s))
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(dp); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				continue
			}
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				pairs := dp[i-1] + 1
				preIndex := i - 2*pairs
				if preIndex > 0 {
					dp[i] = dp[preIndex] + pairs
				} else {
					dp[i] = pairs
				}
			}
		}
	}
	return 2 * max(dp)
}

func max(arr []int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
		}
	}
	return m
}
