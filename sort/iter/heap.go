package iter

import "sync"

func HeapSort(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	heapify(arr, step, wg)
	for i := len(arr) - 1; i > 0; i-- {
		<-step
		arr[i] = deleteMax(arr[:i+1], step, wg) // deleteMaxの中でwg.Done()が呼ばれる
	}

	done <- struct{}{}
}

func heapify(arr []int, c chan struct{}, wg *sync.WaitGroup) {
	for i := len(arr)/2 - 1; i >= 0; i-- {

		downMax(i, arr, len(arr), c, wg)
	}
}

func downMax(i int, arr []int, n int, c chan struct{}, wg *sync.WaitGroup) {
	j := 2*i + 1

	if j >= n {
		return
	}

	if j+1 < n && arr[j] < arr[j+1] {
		j = j + 1
	}

	if arr[j] > arr[i] {

		<-c
		swap(&arr[i], &arr[j])
		wg.Done()

		downMax(j, arr, n, c, wg)
	}

}

func deleteMax(arr []int, c chan struct{}, wg *sync.WaitGroup) int {
	wg.Done()

	max := arr[0]

	<-c
	arr[0] = arr[len(arr)-1]
	wg.Done()

	downMax(0, arr, len(arr)-1, c, wg)
	return max
}
