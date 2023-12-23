package sort

import (
	"sync"
)

const MAX = 1000

func HeapSort(arr []int) {
	heapify(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = deleteMax(arr[:i+1])
	}
}

func HeapSortAsync(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()

	heapifyAsync(arr, step, wg)
	for i := len(arr) - 1; i > 0; i-- {
		<-step
		arr[i] = deleteMaxAsync(arr[:i+1], step, wg) // deleteMaxAsyncの中でwg.Done()が呼ばれる
	}

	done <- struct{}{}
}

func heapify(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		downMax(i, arr, len(arr))
	}
}

func downMax(i int, arr []int, n int) {
	j := 2*i + 1

	if j >= n {
		return
	}

	if j+1 < n && arr[j] < arr[j+1] {
		j = j + 1
	}

	if arr[j] > arr[i] {
		swap(&arr[i], &arr[j])
		downMax(j, arr, n)
	}
}

func deleteMax(arr []int) int {
	max := arr[0]
	arr[0] = arr[len(arr)-1]
	downMax(0, arr, len(arr)-1)
	return max
}

func heapifyAsync(arr []int, c chan struct{}, wg *sync.WaitGroup) {
	for i := len(arr)/2 - 1; i >= 0; i-- {

		downMaxAsync(i, arr, len(arr), c, wg)
	}
}

func downMaxAsync(i int, arr []int, n int, c chan struct{}, wg *sync.WaitGroup) {
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

		downMaxAsync(j, arr, n, c, wg)
	}

}

func deleteMaxAsync(arr []int, c chan struct{}, wg *sync.WaitGroup) int {
	wg.Done()

	max := arr[0]

	<-c
	arr[0] = arr[len(arr)-1]
	wg.Done()

	downMaxAsync(0, arr, len(arr)-1, c, wg)
	return max
}
