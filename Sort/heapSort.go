package Sort

func HeapSort(arr []int) {
	//构建大根堆
	for i:=len(arr)/2-1;i>=0;i-- {
		down(arr,i,len(arr))
	}
	for i:=len(arr)-1;i>=1;i-- {
		swap(arr,0,i)
		down(arr,0,i)
	}
}


func down(arr []int,i, n int) {
	for {
		//左孩子
		j1 := i*2+1
		if j1>=n || j1 < 0 {
			break
		}
		j := j1
		if j2 := i*2+2;j2<n && arr[j2] > arr[j1] {
			j = j2
		}
		//父节点比两个子节点中最大的值还要大，则调整结束
		if arr[i] > arr[j] {
			break
		}
		swap(arr,i,j)
		i = j
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

