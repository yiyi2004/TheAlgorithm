package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	index := len(nums1) - 1
	index1 := m - 1
	index2 := n - 1
	for index >= 0 && index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[index] = nums1[index1]
			index1--
		} else {
			nums1[index] = nums2[index2]
			index2--
		}
		index--
	}
	if index1 < 0 {
		for i := index2; i >= 0; i-- {
			nums1[index] = nums2[i]
			index--
		}
	} else if index2 < 0 {
		for i := index1; i >= 0; i-- {
			nums1[index] = nums1[i]
			index--
		}
	}
}
