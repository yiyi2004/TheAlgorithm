package main

func main() {

}

// 最大值的风险是如果你进行了加的操作可能会导致溢出的
func jump(nums []int) int {
	count := 0
	target := len(nums) - 1
	index := -1
	for target > 0 {
		for i := target; i >= 0; i-- {
			if i+nums[i] >= target {
				index = i
			}
		}
		count++
		target = index
	}
	return count
}
