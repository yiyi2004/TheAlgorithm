package main

import "fmt"

func main() {
	s := "aab"
	res := partition(s)
	fmt.Println(res)
}

// 回溯 + 动态规划预处理
//func partition(s string) (ans [][]string) {
//	n := len(s)
//	f := make([][]bool, n)
//	for i := range f {
//		f[i] = make([]bool, n)
//		for j := range f[i] {
//			f[i][j] = true
//		}
//	}
//	for i := n - 1; i >= 0; i-- {
//		for j := i + 1; j < n; j++ {
//			f[i][j] = s[i] == s[j] && f[i+1][j-1]
//		}
//	}
//
//	splits := []string{}
//	var dfs func(int)
//	dfs = func(i int) {
//		if i == n {
//			ans = append(ans, append([]string(nil), splits...))
//			return
//		}
//		for j := i; j < n; j++ {
//			if f[i][j] {
//				splits = append(splits, s[i:j+1])
//				dfs(j + 1)
//				splits = splits[:len(splits)-1]
//			}
//		}
//	}
//	dfs(0)
//	return
//}

func partition(s string) (res [][]string) {
	if len(s) == 0 {
		return nil
	}
	tmp := make([]string, 0)
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if startIndex == len(s) {
			res = append(res, append([]string{}, tmp...))
			return
		}
		for i := startIndex; i < len(s); i++ {
			str := s[startIndex : i+1]
			if isPalindrome(str) {
				tmp = append(tmp, str)
				backtracking(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtracking(0)
	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
