package main

func main() {

}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func PathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return count(root, sum) + PathSum(root.left, sum) + PathSum(root.right, sum)
}

func count(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	flag := 0
	if root.val == sum {
		flag = 1
	}
	leftCount := count(root.left, sum-root.val)
	rightCount := count(root.right, sum-root.val)

	return flag + leftCount + rightCount
}

//// 你这个求的是子树的和，不是路径和
//func process(root *TreeNode, num int, sum int) int {
//	if root == nil {
//		return 0
//	}
//	leftSum := process(root.left, num, sum)
//	rightSum := process(root.right, num, sum)
//	if leftSum+rightSum+root.val == num {
//		sum++
//	}
//	return leftSum + rightSum + root.val
//}

//int pathSum(TreeNode *root, int sum) {
//if (root == nullptr) return 0;
//return count(root, sum) + pathSum(root->left, sum) +
//pathSum(root->right, sum);
//}
//
//int count(TreeNode *node, int sum) {
//if (node == nullptr) return 0;
//return (node->val == sum) + count(node->left, sum - node->val) +
//count(node->right, sum - node->val);
//}
