package main

import (
	"fmt"
)

func main() {
	s := "abab"
	p := "ab"
	anagrams := findAnagrams(s, p)
	fmt.Println(anagrams)
}

// 使用有限制的数组，而不是 map
// 数组更加节省空间
// 这种方式确实能够 AC，但感觉不是最优解
func findAnagrams(s string, p string) []int {
	pm := make([]int, 26)

	for i := 0; i < len(p); i++ {
		pm[p[i]-'a']++
	}
	var res []int

	left := 0
	right := left + len(p) - 1
	for right < len(s) {
		sm := make([]int, 26)
		for i := left; i <= right; i++ {
			sm[s[i]-'a']++
		}
		if isSameMap(sm, pm) {
			res = append(res, left)
		}
		left++
		right++
	}

	return res
}

func isSameMap(m1, m2 []int) bool {
	for i := 0; i < len(m2); i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}
