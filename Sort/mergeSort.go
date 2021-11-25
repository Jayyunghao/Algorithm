package Sort

func MergeSort(arr []int) {
	mergeSortCore(arr,0,len(arr)-1)
	return
}

func mergeSortCore(arr []int, start, end int) {
	if start < end {
		mid := (start+end) >> 1
		mergeSortCore(arr,start,mid)
		mergeSortCore(arr,mid+1,end)
		merge(arr,start,mid,end)
	}
}

func merge(arr []int,start, mid, end int) {
	temp := make([]int, end-start+1)
	index := 0
	i,j := start, mid+1
	for i<=mid && j <= end {
		if arr[i] < arr[j] {
			temp[index] = arr[i]
			index++
			i++
		} else {
			temp[index] = arr[j]
			j++
			index ++
		}
	}
	for;i<=mid;i++ {
		temp[index] = arr[i]
		index ++
	}
	for ;j<= end;j++ {
		temp[index] = arr[j]
		index ++
	}
	for k:=0;k<index;k++ {
		arr[start+k] = temp[k]
	}
}
