package main

func main() {

}

func removeElement(nums []int, val int) int {
	p, q := 0, 0
	for p < len(nums) && nums[p] != val {
		p++
	}
	q = p + 1
	for q < len(nums) {
		if nums[q] == val {
			q++
			continue
		}
		nums[p] = nums[q]
		p++
		q++
	}
	return p
}
