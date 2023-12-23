package iter

import "sync"

func QuickSort(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	quickSort(arr, 0, len(arr)-1, step, wg)

	done <- struct{}{}
}

func quickSort(arr []int, left, right int, c chan struct{}, wg *sync.WaitGroup) {
	if left >= right {
		return
	}

	pivot := partition(arr, left, right, c, wg)
	quickSort(arr, left, pivot-1, c, wg)
	quickSort(arr, pivot+1, right, c, wg)
}

func partition(arr []int, left, right int, c chan struct{}, wg *sync.WaitGroup) int {
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
