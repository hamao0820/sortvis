package iter

import "sync"

func ShellSort(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	var i, j int

	H := [5]int{57, 23, 10, 4, 1}

	for _, h := range H {
		for i = h; i < len(arr); i++ {
			j = i
			for j >= h && arr[j] < arr[j-h] {
				<-step
				swap(&arr[j], &arr[j-h])
				wg.Done()
				j -= h
			}
		}
	}

	done <- struct{}{}
}
