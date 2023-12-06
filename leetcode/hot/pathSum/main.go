package main

import (
	"fmt"
)

func main() {
	n1 := new(TreeNode)
	n2 := new(TreeNode)
	n3 := new(TreeNode)
	n4 := new(TreeNode)
	n5 := new(TreeNode)
	n1.Val = 1
	n2.Val = 2
	n3.Val = 3
	n4.Val = 4
	n5.Val = 5
	n1.Right = n2
	n2.Right = n3
	n3.Right = n4
	n4.Right = n5
	res := pathSum(n1, 3)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func rootSum(root *TreeNode, targetSum int) (res int) {
//	if root == nil {
//		return
//	}
//	val := root.Val
//	if val == targetSum {
//		res++
//	}
//	res += rootSum(root.Left, targetSum-val)
//	res += rootSum(root.Right, targetSum-val)
//	return
//}
//
//func pathSum(root *TreeNode, targetSum int) int {
//	if root == nil {
//		return 0
//	}
//	res := rootSum(root, targetSum)
//	res += pathSum(root.Left, targetSum)
//	res += pathSum(root.Right, targetSum)
//	return res
//}

// 我不知道哪里出了问题
//
//	type info struct {
//		pathNum int
//	}
//func pathSum(root *TreeNode, targetSum int) int {
//	if root == nil {
//		return 0
//	}
//	res := 0
//	res += rootSum(root, targetSum, 0)
//	res += pathSum(root.Left, targetSum)
//	res += pathSum(root.Right, targetSum)
//	return res
//}
//
//func rootSum(root *TreeNode, targetSum int, preVal int) (res int) {
//	if root == nil {
//		return
//	}
//	if root.Val+preVal == targetSum {
//		res++
//	}
//	res += rootSum(root.Left, targetSum, preVal+root.Val)
//	res += rootSum(root.Right, targetSum, preVal+root.Val)
//	return res
//}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return rootSum(root, targetSum, 0) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func rootSum(root *TreeNode, targetSum int, preVal int) int {
	if root == nil {
		return 0
	}
	res := 0
	if preVal+root.Val == targetSum {
		res++
	}
	return res + rootSum(root.Left, targetSum, preVal+root.Val) + rootSum(root.Right, targetSum, preVal+root.Val)
}
