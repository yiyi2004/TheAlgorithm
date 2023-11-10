package main

func main() {

}

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	n := len(nums)
	var res [][]int
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, path...))
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return res
}
