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

	dp := make([]int, len(s))
	stack := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, "(")
		} else {
			if len(stack) == 0 {
				continue
			}
			if stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
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
	fmt.Println(dp)
	return max(dp) * 2
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
