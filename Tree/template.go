package Tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(x int) *TreeNode {
	return &TreeNode{
		Val: x,
		Left: nil,
		Right: nil,
	}
}

//前序递归遍历
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ",root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

//前序非递归遍历
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int,0)
	stack := make([]*TreeNode,0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result,root.Val)
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

//中序遍历
func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Printf("%d ",root.Val)
	InOrder(root.Right)
}

//非递归中序遍历
func InOrderTraversal(root *TreeNode) []int{
	if root == nil {
		return nil
	}
	result := make([]int,0)
	stack := make([]*TreeNode,0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result,node.Val)
		root = node.Right
	}
	return result
}

//后序遍历
func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d ",root.Val)
}

//非递归后续遍历
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int,0)
	stack := make([]*TreeNode,0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result,node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}
