package main

func main() {

}

func maxProfit(prices []int) int {
	left := 0
	total := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[left] {
			left = i
		} else {
			if i+1 < len(prices) && prices[i+1] > prices[i] {
				continue
			} else {
				total += prices[i] - prices[left]
				left = i + 1
			}
		}
	}

	return total
}
