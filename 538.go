func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}

	traverse(root)

	return root
}
