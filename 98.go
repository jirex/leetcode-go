package main

import "math"

func main() {

}

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(root *TreeNode, low, up int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= up {
		return false
	}

	return check(root.Left, low, root.Val) && check(root.Right, root.Val, up)
}
