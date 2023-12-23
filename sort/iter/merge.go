package iter

import "sync"

func MergeSort(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	mergeSort(arr, 0, len(arr)-1, step, wg)

	done <- struct{}{}
}

func mergeSort(arr []int, left, right int, c chan struct{}, wg *sync.WaitGroup) {
	if left < right {
		mid := (left + right) / 2

		mergeSort(arr, left, mid, c, wg)
		mergeSort(arr, mid+1, right, c, wg)

		merge(arr, left, mid, right, c, wg)
	}
}

func merge(arr []int, left, mid, right int, c chan struct{}, wg *sync.WaitGroup) {
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
		<-c
		arr[left+i] = tmp[i]
		wg.Done()
	}
}
