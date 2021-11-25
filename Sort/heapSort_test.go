package Sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := make([]int,10)
	for i:=9;i>=1;i-- {
		arr[9-i] = i
	}
	HeapSort(arr)
	for i:=0;i<len(arr);i++ {
		fmt.Printf("%d",arr[i])
	}
}
