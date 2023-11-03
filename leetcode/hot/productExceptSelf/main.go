package main

func main() {

}

func productExceptSelf(nums []int) []int {
	containTwoZero := 0
	exceptZeroProduct := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			containTwoZero++
			continue
		} else {
			exceptZeroProduct *= nums[i]
		}
	}

	if containTwoZero > 1 {
		for i := 0; i < len(nums); i++ {
			nums[i] = 0
		}
	} else if containTwoZero == 1 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				nums[i] = exceptZeroProduct
			} else {
				nums[i] = 0
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			nums[i] = exceptZeroProduct / nums[i]
		}
	}
	return nums
}
