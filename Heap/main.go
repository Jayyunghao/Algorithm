package main

import (
	"fmt"
	"log"
	"math"
)

type MaxHeap struct {
	num []int
}

func constructor (nums []int) *MaxHeap {
	array := []int{math.MaxInt32} //哨兵
	array = append(array,nums...)
	//向下调整
	n := len(array)-1
	for i := n/2;i>0;i-- {
		down(array,i,n)
	}
	h := &MaxHeap{num: array}
	return h
}

func(m *MaxHeap) Push(x int) {
	m.num = append(m.num,x)
	i := len(m.num)-1
	//上浮
	up(m.num,i)
}

func(m *MaxHeap) Pop() (int,error) {
	if len(m.num) <= 1 {
		return 0, fmt.Errorf("MaxHeap is empty")
	}
	n := len(m.num)-1
	//调换最大堆根节点到最后，然后进行调整
	m.num[1], m.num[n] = m.num[n], m.num[1]
	down(m.num,1,n-1)
	maxElement := m.num[n]
	m.num = m.num[:n]
	return maxElement, nil
}

func up(num []int, j int) {
	for {
		//父节点
		i := j/2
		if i < 1 || num[j] < num[i] {
			break
		}
		//子节点比父节点大，则替换
		num[i],num[j] = num[j], num[i]
		j = i
	}
}

func down(num []int, i, n int) {
	for {
		//孩子节点
		j1 := 2*i
		if j1 >= n || j1 <1 {
			break   //0 after int overflow
		}
		j := j1
		if j2 := j1+1 ;j2<=n && num[j2] > num[j1] {
			j = j2
		}
		//父节点比两个子节点中最大的值还要大，则调整结束
		if num[i] > num[j] {
			break
		}
		num[i],num[j] = num[j],num[i]
		i = j
	}
}

func main() {
	maxHeap := constructor([]int{3,7,8})
	maxHeap.Push(10)
	maxHeap.Push(1)
	maxHeap.Push(11)
	n,err := maxHeap.Pop()
	if err != nil {
		log.Fatalf("pop err:%s",err.Error())
	}
	fmt.Println(n)
}
