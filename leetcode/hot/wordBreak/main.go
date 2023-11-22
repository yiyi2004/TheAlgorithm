package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	res := wordBreak(s, wordDict)
	fmt.Println(res)
}

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]struct{})
	for _, w := range wordDict {
		wordDictSet[w] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && exists(wordDictSet, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func exists(help map[string]struct{}, v string) bool {
	_, ok := help[v]
	return ok
}
