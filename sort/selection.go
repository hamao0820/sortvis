package sort

import "sync"

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		swap(&arr[i], &arr[min])
	}
}

func SelectionSortAsync(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		<-step
		swap(&arr[i], &arr[min])
		wg.Done()
	}

	done <- struct{}{}
}
