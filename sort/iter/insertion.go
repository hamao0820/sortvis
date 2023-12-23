package iter

import "sync"

func InsertionSort(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()
	var i, j int

	for i = 1; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j] < arr[j-1] {
			<-step
			swap(&arr[j], &arr[j-1])
			wg.Done()
			j--
		}
	}

	done <- struct{}{}
}
