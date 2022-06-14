package main

import "strconv"

func main() {
	

}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	mp := map[string]int{}
	strMp := map[string]*TreeNode{}
	var dfs func(*TreeNode) string
	dfs  = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		var str string
		str += "."+strconv.Itoa(root.Val)
		str += dfs(root.Left)
		str += dfs(root.Right)
		mp[str]++
		strMp[str] = root
		return str
	}
	dfs(root)

	res := []*TreeNode{}
	for i := range mp {
		if mp[i] > 1 {
			res = append(res, strMp[i])
		}
	}
	return res
}
