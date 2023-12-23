package sort

import "sync"

func QuickSortAsync(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	quickSortAsync(arr, 0, len(arr)-1, step, wg)

	done <- struct{}{}
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			swap(&arr[i], &arr[j])
			i++
		}
	}
	swap(&arr[i], &arr[right])
	return i
}

func quickSortAsync(arr []int, left, right int, c chan struct{}, wg *sync.WaitGroup) {
	if left >= right {
		return
	}

	pivot := partitionAsync(arr, left, right, c, wg)
	quickSortAsync(arr, left, pivot-1, c, wg)
	quickSortAsync(arr, pivot+1, right, c, wg)
}

func partitionAsync(arr []int, left, right int, c chan struct{}, wg *sync.WaitGroup) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			<-c
			arr[i], arr[j] = arr[j], arr[i]
			wg.Done()
			i++
		}
	}
	<-c
	swap(&arr[i], &arr[right])
	wg.Done()
	return i
}
