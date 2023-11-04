package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 同故宫是通过了，但是时间和空间复杂度都比较低
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	root := new(TreeNode)
	leftL, leftR, mid, rightL, rightR := 0, 0, 0, 0, 0
	root.Val = preorder[0]
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			mid = i
			break
		}
	}
	leftL, leftR, rightL, rightR = 0, mid-1, mid+1, len(inorder)-1

	return createTree(preorder, inorder, root, leftL, leftR, rightL, rightR)
}

func createTree(preorder []int, inorder []int, root *TreeNode, leftL, leftR, rightL, rightR int) *TreeNode {
	root.Left = createLeftTree(preorder, inorder, leftL, leftR, 0+1)
	root.Right = createRightTree(preorder, inorder, rightL, rightR, 0+leftR-leftL+2)
	return root
}

func createLeftTree(preorder []int, inorder []int, leftL, leftR int, rootIndex int) *TreeNode {
	if leftR < 0 || leftL > leftR {
		return nil
	}
	if leftR-leftL == 0 {
		node := TreeNode{}
		node.Val = inorder[leftL]
		return &node
	}
	root := new(TreeNode)
	for i := leftL; i <= leftR; i++ {
		if inorder[i] == preorder[rootIndex] {
			root.Val = inorder[i]
			root.Left = createLeftTree(preorder, inorder, leftL, i-1, rootIndex+1)
			root.Right = createRightTree(preorder, inorder, i+1, leftR, rootIndex+1+i-leftL)
			break
		}
	}

	return root
}

func createRightTree(preorder []int, inorder []int, rightL, rightR int, rootIndex int) *TreeNode {
	if rightL >= len(inorder) || rightL > rightR {
		return nil
	}
	if rightR-rightL == 0 {
		node := TreeNode{}
		node.Val = inorder[rightL]
		return &node
	}
	root := new(TreeNode)
	for i := rightL; i <= rightR; i++ {
		if inorder[i] == preorder[rootIndex] {
			root.Val = inorder[i]
			root.Left = createLeftTree(preorder, inorder, rightL, i-1, rootIndex+1)
			root.Right = createRightTree(preorder, inorder, i+1, rightR, rootIndex+1+i-rightL)
			break
		}
	}

	return root
}
