package main

func main() {
	candy([]int{1, 2, 87, 87, 87, 2, 1})
}

func candy(ratings []int) int {
	left := make([]int, len(ratings))

	for i, rating := range ratings {
		if i > 0 && rating > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := make([]int, len(ratings))
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}
	res := 0
	for i := 0; i < len(left); i++ {
		res += max(left[i], right[i])
	}
	return res
}

func max(arr ...int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
