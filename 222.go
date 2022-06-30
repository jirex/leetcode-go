func countNodes(root *TreeNode) int {
	l := root
	r := root

	hl, hr := 0, 0

	for l != nil {
		l = l.Left
		hl++
	}

	for r != nil {
		r = r.Right
		hr++
	}

	if hl == hr {
		return int(math.Pow(float64(2), float64(hl)) - 1)
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
