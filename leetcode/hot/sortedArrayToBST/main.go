package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	root := sortedArrayToBST(nums)
	printTree(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) {
	for root == nil {
		return
	}

	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	return createBST(nums, 0, len(nums)-1)
}

func createBST(nums []int, l, r int) *TreeNode {
	if r < l {
		return nil
	}
	if r-l == 0 {
		return &TreeNode{
			Val:   nums[l],
			Left:  nil,
			Right: nil,
		}
	}
	mid := (l + r) >> 1
	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = createBST(nums, l, mid-1)
	root.Right = createBST(nums, mid+1, r)
	return root
}
