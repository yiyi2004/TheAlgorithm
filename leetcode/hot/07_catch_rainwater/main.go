package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(arr)
	fmt.Println(res)
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	preMax := make([]int, len(height))
	postMax := make([]int, len(height))

	preMax[0] = height[0]
	postMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height)-1; i++ {
		// initialize the preMax
		if height[i] > preMax[i-1] {
			preMax[i] = height[i]
		} else {
			preMax[i] = preMax[i-1]
		}

	}
	for i := len(height) - 2; i >= 0; i-- {
		// initialize the postMax
		if height[i] > postMax[i+1] {
			postMax[i] = height[i]
		} else {
			postMax[i] = postMax[i+1]
		}
	}

	sum := 0
	for i := 0; i < len(height)-1; i++ {
		sum += min(preMax[i], postMax[i]) - height[i]
	}
	return sum
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}
