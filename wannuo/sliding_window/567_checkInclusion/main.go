package main

func main() {
	checkInclusion("ab", "eidbaooo")
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Char := [26]int{}
	s2Char := [26]int{}
	for i := 0; i < len(s1); i++ {
		s1Char[s1[i]-'a']++
	}
	for i := 0; i < len(s1); i++ {
		s2Char[s2[i]-'a']++
	}
	if check(s1Char, s2Char) {
		return true
	}
	for l, r := 0, len(s1); r < len(s2); {
		s2Char[s2[r]-'a']++
		s2Char[s2[l]-'a']--
		if check(s1Char, s2Char) {
			return true
		}
		l++
		r++
	}

	return false
}

func check(s1Char, s2Char [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1Char[i] != s2Char[i] {
			return false
		}
	}
	return true
}
