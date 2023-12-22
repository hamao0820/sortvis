package sort

import (
	"sync"
)

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func BubbleSortAsync(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				<-step
				swap(&arr[j], &arr[j+1])
				wg.Done()
			}
		}
	}

	done <- struct{}{}
}
