package main

import (
	"fmt"
	"math"
)

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

//func maxPathSum(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return process(root).maxVal
//}

//type info struct {
//	deepMax int
//	maxVal  int
//}
//
//func process(root *TreeNode) *info {
//	if root == nil {
//		return nil
//	}
//	leftInfo := process(root.Left)
//	rightInfo := process(root.Right)
//	res := &info{}
//	if leftInfo == nil && rightInfo == nil {
//		res.maxVal = root.Val
//		res.deepMax = root.Val
//	} else if leftInfo != nil && rightInfo == nil {
//		res.maxVal = max(root.Val, root.Val+leftInfo.deepMax, leftInfo.maxVal)
//		res.deepMax = max(root.Val, root.Val+leftInfo.deepMax)
//	} else if leftInfo == nil && rightInfo != nil {
//		res.maxVal = max(root.Val, root.Val+rightInfo.deepMax, rightInfo.maxVal)
//		res.deepMax = max(root.Val, root.Val+rightInfo.deepMax)
//	} else {
//		res.maxVal = max(root.Val, root.Val+leftInfo.deepMax, root.Val+rightInfo.deepMax, root.Val+leftInfo.deepMax+rightInfo.deepMax, leftInfo.maxVal, rightInfo.maxVal)
//		res.deepMax = max(root.Val, root.Val+leftInfo.deepMax, root.Val+rightInfo.deepMax)
//	}
//	return res
//}
//
//func max(nums ...int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	maxVal := nums[0]
//	for _, v := range nums {
//		if v > maxVal {
//			maxVal = v
//		}
//	}
//	return maxVal
//}

type info struct {
	maxPathSum       int
	maxSinglePathSum int
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return process(root).maxPathSum
}

func process(root *TreeNode) *info {
	if root == nil {
		return nil
	}
	res := info{
		maxPathSum:       0,
		maxSinglePathSum: 0,
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)

	if leftInfo != nil && rightInfo != nil {
		res.maxSinglePathSum = max(root.Val, leftInfo.maxSinglePathSum+root.Val, rightInfo.maxSinglePathSum+root.Val)
		res.maxPathSum = max(root.Val, leftInfo.maxPathSum, rightInfo.maxPathSum, root.Val+leftInfo.maxSinglePathSum+rightInfo.maxSinglePathSum, root.Val+leftInfo.maxSinglePathSum, root.Val+rightInfo.maxSinglePathSum)
	} else if leftInfo != nil && rightInfo == nil {
		res.maxSinglePathSum = max(root.Val, leftInfo.maxSinglePathSum+root.Val)
		res.maxPathSum = max(root.Val, leftInfo.maxPathSum, root.Val+leftInfo.maxSinglePathSum)
	} else if leftInfo == nil && rightInfo != nil {
		res.maxSinglePathSum = max(root.Val, rightInfo.maxSinglePathSum+root.Val)
		res.maxPathSum = max(root.Val, rightInfo.maxPathSum, root.Val+rightInfo.maxSinglePathSum)
	} else {
		res.maxPathSum = root.Val
		res.maxSinglePathSum = root.Val
	}

	return &res
}

func max(vals ...int) int {
	maxVal := math.MinInt
	for i := 0; i < len(vals); i++ {
		if vals[i] > maxVal {
			maxVal = vals[i]
		}
	}
	return maxVal
}
