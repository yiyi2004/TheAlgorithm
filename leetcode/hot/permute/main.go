package main

func main() {

}

//func permute(nums []int) [][]int {
//	if nums == nil {
//		return nil
//	}
//	n := len(nums)
//	var res [][]int
//	path := make([]int, n)
//	onPath := make([]bool, n)
//	var dfs func(i int)
//	dfs = func(i int) {
//		if i == n {
//			res = append(res, append([]int{}, path...))
//		}
//		for j, on := range onPath {
//			if !on {
//				path[i] = nums[j]
//				onPath[j] = true
//				dfs(i + 1)
//				onPath[j] = false
//			}
//		}
//	}
//	dfs(0)
//	return res
//}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0)
	path := make([]int, len(nums))
	onPath := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		onPath[i] = false
	}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			// 为什么需要这么写呢？
			// 创建当前的数组，而不是引用原始数组
			// 原始数组的数据是最后一次递归得到的结果
			// 所以得到的 res 中的元素都是相同的
			res = append(res, append([]int{}, path...))
			return
		}
		for k, ok := range onPath {
			if !ok {
				path[index] = nums[k]
				onPath[k] = true
				dfs(index + 1)
				onPath[k] = false
			}
		}
	}
	dfs(0)
	return res
}
