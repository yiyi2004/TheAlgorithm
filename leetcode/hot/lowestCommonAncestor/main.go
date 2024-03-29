package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type info struct {
}

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root == nil || p == nil || q == nil {
//		return nil
//	}
//	if root == p || root == q {
//		return root
//	}
//	left := lowestCommonAncestor(root.Left, p, q)
//	right := lowestCommonAncestor(root.Right, p, q)
//	if left != nil && right != nil {
//		return root
//	}
//	if left == nil {
//		return right
//	}
//	return left
//}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
