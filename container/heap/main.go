package main

import (
	"container/heap"
	"fmt"
)

//最大堆
type MaxHeap []int

func(h MaxHeap) Len() int {return len(h)}
//An MaxHeap is a max-heap of ints
func(h MaxHeap) Less(i,j int)bool {return h[i]>h[j]}
func(h MaxHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}

func(h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func(h *MaxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0:n-1]
	return x
}
func main() {
	h := &MaxHeap{2,10,6}
	heap.Init(h)
	heap.Push(h,5)
	n := heap.Pop(h)
	fmt.Println(n)
}
