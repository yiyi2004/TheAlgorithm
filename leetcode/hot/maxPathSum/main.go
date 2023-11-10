package main

import "fmt"

func main() {
	root := new(TreeNode)
	left := new(TreeNode)
	right := new(TreeNode)
	root.Val = 1
	left.Val = -2
	right.Val = 3
	root.Left = left
	root.Right = right
	res := maxPathSum(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return process(root).maxVal
}

type info struct {
	deepMax int
	maxVal  int
}

func process(root *TreeNode) *info {
	if root == nil {
		return nil
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)
	res := &info{}
	if leftInfo == nil && rightInfo == nil {
		res.maxVal = root.Val
		res.deepMax = root.Val
	} else if leftInfo != nil && rightInfo == nil {
		res.maxVal = max(root.Val, root.Val+leftInfo.deepMax, leftInfo.maxVal)
		res.deepMax = max(root.Val, root.Val+leftInfo.deepMax)
	} else if leftInfo == nil && rightInfo != nil {
		res.maxVal = max(root.Val, root.Val+rightInfo.deepMax, rightInfo.maxVal)
		res.deepMax = max(root.Val, root.Val+rightInfo.deepMax)
	} else {
		res.maxVal = max(root.Val, root.Val+leftInfo.deepMax, root.Val+rightInfo.deepMax, root.Val+leftInfo.deepMax+rightInfo.deepMax, leftInfo.maxVal, rightInfo.maxVal)
		res.deepMax = max(root.Val, root.Val+leftInfo.deepMax, root.Val+rightInfo.deepMax)
	}
	return res
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
