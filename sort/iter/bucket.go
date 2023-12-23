package iter

import "sync"

func BucketSort(arr []int, step, done chan struct{}, wg *sync.WaitGroup) {
	wg.Done()
	defer (func() { done <- struct{}{} })()

	n := len(arr)
	if n <= 1 {
		return
	}

	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	bucket := make([]int, max-min+1)
	for _, v := range arr {
		bucket[v-min]++
	}

	i := 0
	for j := 0; j < len(bucket); j++ {
		for bucket[j] > 0 {
			<-step
			arr[i] = j + min
			wg.Done()
			i++
			bucket[j]--
		}
	}
}
