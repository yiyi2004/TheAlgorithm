package main

func main() {

}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		if len(path) >= 2 {
			res = append(res, append([]int{}, path...))
		}
		help := make(map[int]struct{})
		for i := index; i < len(nums); i++ {
			if len(path) >= 1 && nums[i] < path[len(path)-1] {
				continue
			}
			if _, ok := help[nums[i]]; ok {
				continue
			}
			help[nums[i]] = struct{}{}
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return res
}
