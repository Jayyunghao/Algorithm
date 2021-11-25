package Sort


func QuickSort(arr []int) {
	QuickSortCore(arr,0,len(arr)-1)
	return
}

func QuickSortCore(arr []int,start, end int) {
	if start < end {
		privot := partition(arr, start, end)
		QuickSortCore(arr,start,privot-1)
		QuickSortCore(arr,privot+1,end)
	}
}

func partition(arr []int, start, end int) int {
	privot := arr[start]
	i := start +1
	for j:=start+1;j<=end;j++ {
		if arr[j]<privot {
			swap(arr,i,j)
			i++
		}
	}
	swap(arr,i-1,start)
	return i-1
}