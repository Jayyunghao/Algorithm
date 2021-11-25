package SortList

import (
	"fmt"
	"testing"
)

func TestQuickSortList2(t *testing.T) {
	headPre := new(ListNode)
	curPre := headPre
	for i:=10;i>=0;i-- {
		curPre.Next = new(ListNode)
		curPre.Next.Val = i
		curPre = curPre.Next
	}
	curPre.Next = nil
	head := headPre.Next
	headPre.Next = QuickSortList2(head)
	for cur := headPre.Next;cur != nil ;cur = cur.Next {
		fmt.Println(cur.Val)
	}
}

func TestQuickSortList(t *testing.T) {
	headPre := new(ListNode)
	curPre := headPre
	for i:=10;i>=0;i-- {
		curPre.Next = new(ListNode)
		curPre.Next.Val = i
		curPre = curPre.Next
	}
	curPre.Next = nil
	head := headPre.Next
	headPre.Next = QuickSortList(head)
	for cur := headPre.Next;cur != nil ;cur = cur.Next {
		fmt.Println(cur.Val)
	}
}
