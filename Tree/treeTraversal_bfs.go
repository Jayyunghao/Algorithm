package Tree

//bfs  典型场景层序遍历
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int,0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	for len(queue) != 0 {
		list := make([]int,0)
		sz := len(queue)
		for i:=0;i<sz;i++ {
			cur := queue[0]
			queue = queue[1:]
			list = append(list,cur.Val)
			if cur.Left != nil {
				queue = append(queue,cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue,cur.Right)
			}
		}
		result = append(result,list)
	}
	return result
}
