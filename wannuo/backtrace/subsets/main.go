package main

func main() {

}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		res = append(res, append([]int{}, path...))

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return res
}
