package Tree

import (
	"strconv"
	"testing"
)

func TestPreOrder(t *testing.T) {
	root := NewTreeNode(8)
	left := NewTreeNode(5)
	right := NewTreeNode(7)
	lleft := NewTreeNode(6)
	lRight := NewTreeNode(10)
	rLeft := NewTreeNode(11)
	left.Left = lleft
	left.Right = lRight
	right.Left = rLeft
	root.Left = left
	root.Right = right
	//PreOrder(root)
	preCheck := "8 5 6 10 7 11 "
	preOrder := PreOrderTraversal(root)
	if len(preOrder) == 0 {
		t.Error("前序非递归遍历失败")
	}
	var s string
	for i:=0;i<len(preOrder);i++ {
		s += strconv.Itoa(preOrder[i]) + " "
	}
	if preCheck != s {
		t.Errorf("前序遍历错误,遍历顺序:%s",s)
	}

	//InOrder(root)
	InCheck := "6 5 10 8 11 7 "
	inOrder := InOrderTraversal(root)
	if len(inOrder) == 0 {
		t.Error("中序非递归遍历失败")
	}
	var s1 string
	for i:=0;i<len(inOrder);i++ {
		s1 += strconv.Itoa(inOrder[i]) + " "
	}
	if InCheck != s1 {
		t.Errorf("中序遍历错误,遍历顺序:%s",s1)
	}

	PostOrder(root)
	postCheck := "6 10 5 11 7 8 "
	postOrder := PostOrderTraversal(root)
	var s2 string
	for i:=0;i<len(postOrder);i++ {
		s2 += strconv.Itoa(postOrder[i]) + " "
	}
	if postCheck != s2 {
		t.Errorf("后序遍历错误,遍历顺序:%s",s2)
	}
}