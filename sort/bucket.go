package sort

func BucketSort(arr []int) {
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
			arr[i] = j + min
			i++
			bucket[j]--
		}
	}
}

func BucketSortAsync(arr []int, c chan int) {
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
			<-c
			arr[i] = j + min
			i++
			bucket[j]--
		}
	}
}
