package SortList

type ListNode struct {
	Val   	int
	Next   *ListNode
}
/*快排链表*/
func QuickSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {return head}
	QuickSort(head,nil)
	return head
}

func QuickSort(head, tail *ListNode)  {
	if head == tail  || head.Next == tail {return }
	mid := partition(head,tail)
	QuickSort(head,mid.Next)
	QuickSort(mid.Next,tail)
}
//交换节点值
func partition(head, tail *ListNode) *ListNode {
	if head == tail || head.Next == tail {return head}
	privo := head.Val
	pre := head
	p1 := head.Next
	for p2 := head.Next; p2 != tail; {
		if p2.Val < privo {
			temp := p2.Val
			p2.Val = p1.Val
			p1.Val = temp
			pre = p1
			p1 = p1.Next
		}
		p2 = p2.Next
	}
	head.Val = pre.Val
	pre.Val = privo
	return pre
}

func QuickSortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {return head}
	//不改变节点进行排序
	privo := head.Val
	lPre := new(ListNode)
	rPre := new(ListNode)
	l,r, cur := lPre, rPre, head
	for ;cur != nil; {
		next := cur.Next
		if cur.Val < privo {
			l.Next = cur
			l = cur
		} else {
			r.Next = cur
			r = r.Next
		}
		cur = next
	}
	//拼接,现在就是一整个分隔好的链表
	l.Next = rPre.Next
	r.Next = nil
	//先排右边再排左边，有些像归并排序
	right := QuickSortList2(head.Next) //返回右边排序好的部分，头结点
	head.Next = nil                    //将中间的节点跟在左边排序节点后面
	left := QuickSortList2(lPre.Next)
	//再次进行完整拼接
	cur = left
	for ;cur.Next!= nil;cur=cur.Next {}
	cur.Next = right
	return left
}