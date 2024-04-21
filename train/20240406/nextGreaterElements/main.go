package main

func main() {

}

// 单调栈 11min
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, append([]int{}, nums...)...)
	stack := make([]int, 0)
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	stack = append(stack, 0)
	for i := 1; i < len(nums); i++ {
		if nums[stack[len(stack)-1]] >= nums[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
				res[stack[len(stack)-1]] = nums[i]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return res[:n]
}
