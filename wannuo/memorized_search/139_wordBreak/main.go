package main

func main() {
	wordBreak("aaaaaaa", []string{"aaa", "aaaa"})
}

func wordBreak(s string, wordDict []string) bool {
	help := make(map[string]struct{}, len(wordDict))
	for _, v := range wordDict {
		help[v] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			_, ok := help[s[j:i]]
			// dp[j] 是经常的判断田间
			if dp[j] && ok {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
