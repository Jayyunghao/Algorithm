package Tree

//dfs  自上而下(递归)
func dfsOfUpToDown(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result,root.Val)
	dfsOfUpToDown(root.Left,result)
	dfsOfUpToDown(root.Right,result)
}

//dfs 从下到上(分治法)
func dfsOfDownToUp(root *TreeNode) []int {
	result := make([]int,0)
	if root == nil {
		return result
	}
	left := dfsOfDownToUp(root.Left)
	right := dfsOfDownToUp(root.Right)
	result = append(result,root.Val)
	result = append(result,left...)
	result = append(result,right...)
	return result
}
