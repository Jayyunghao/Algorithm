package main

import "fmt"

//双向链表

type Node struct {
	prev *Node
	next *Node
	list *List
	val  int
}

func(n *Node) Next() *Node {
	if p := n.next; p.list != nil && p!= &n.list.root {
		return p
	}
	return nil
}

func(n *Node) Prev() *Node {
	if p := n.prev; p.list != nil && p!= &n.list.root {
		return p
	}
	return nil
}

type List struct {
	root Node
	len int
}

func(l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List {return new(List).Init()}

func(l *List) Len() int {return l.len}

func(l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func(l *List) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insert n after at
func(l *List) insert(n, at *Node) *Node {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
	return n
}

func(l *List) remove(n *Node) *Node {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	n.list = nil
	l.len--
	return n
}

func(l *List) Remove(n *Node) *Node {
	if n.list == l {
		return l.remove(n)
	}
	return nil
}

func(l *List) PushFront(v int) *Node {
	n := &Node{
		prev: nil,
		next: nil,
		list: l,
		val: v,
	}
	return l.insert(n,&l.root)
}

func(l *List) PushBack(v int) *Node {
	n := &Node{
		prev: nil,
		next: nil,
		list: l,
		val: v,
	}
	return l.insert(n,l.root.prev)
}

func main() {
	l := New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	fmt.Println(l.Front().val)
	for e := l.Front();e != nil; e=e.Next() {
		fmt.Println(e.val)
	}
}

