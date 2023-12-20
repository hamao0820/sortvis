package sort

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func MergeSortAsync(arr []int, channel chan int) {
	mergeSortAsync(arr, 0, len(arr)-1, channel)
}

func merge(arr []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i := left
	j := mid + 1
	k := 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = arr[j]
		j++
		k++
	}

	for i := 0; i < k; i++ {
		arr[left+i] = tmp[i]
	}
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2

		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)

		merge(arr, left, mid, right)
	}
}

func mergeSortAsync(arr []int, left, right int, channel chan int) {
	if left < right {
		mid := (left + right) / 2

		mergeSortAsync(arr, left, mid, channel)
		mergeSortAsync(arr, mid+1, right, channel)

		<-channel
		merge(arr, left, mid, right)
	}
}
