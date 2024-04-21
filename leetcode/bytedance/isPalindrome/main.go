package main

func main() {

}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverseHalf := 0
	for x > reverseHalf {
		reverseHalf = reverseHalf*10 + x%10
		x = x / 10
	}

	return x == reverseHalf || x == reverseHalf/10
}
