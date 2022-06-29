func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	val1 := min(p.Val, q.Val)
	val2 := max(p.Val, q.Val)
	return find(root, val1, val2)
}

func find(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val2 {
		return find(root.Left, val1, val2)
	}

	if root.Val < val1 {
		return find(root.Right, val1, val2)
	}

	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
