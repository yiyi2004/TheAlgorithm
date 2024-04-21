package main

func main() {

}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	flag := true
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			for j := 0; j < len(needle); j++ {
				if j+i >= len(haystack) {
					flag = false
					break
				}
				if haystack[j+i] != needle[j] {
					flag = false
					break
				}
			}
			if flag {
				return i
			}
			flag = true
		}
	}
	return -1
}

// KMP 算法是一个快速查找匹配串的算法，它的作用其实就是本题问题：如何快速在「原字符串」中找到「匹配字符串」。
