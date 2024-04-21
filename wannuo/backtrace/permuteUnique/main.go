package main

func main() {

}

// 不好
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(path []int)
	help := make(map[int]struct{})
	dfs = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		repeatHelp := make(map[int]struct{})
		for i := 0; i < len(nums); i++ {
			if _, ok := help[i]; ok {
				continue
			}
			if _, ok := repeatHelp[nums[i]]; ok {
				continue
			}
			repeatHelp[nums[i]] = struct{}{}
			help[i] = struct{}{}
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			delete(help, i)
		}
	}
	dfs([]int{})
	return res
}
