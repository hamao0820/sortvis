package sort

func ShellSort(arr []int) {
	var i, j int

	H := [5]int{57, 23, 10, 4, 1}

	for _, h := range H {
		for i = h; i < len(arr); i++ {
			j = i
			for j >= h && arr[j] < arr[j-h] {
				swap(&arr[j], &arr[j-h])
				j -= h
			}
		}
	}
}

func ShellSortAsync(arr []int, c chan struct{}) {
	var i, j int

	H := [5]int{57, 23, 10, 4, 1}

	for _, h := range H {
		for i = h; i < len(arr); i++ {
			j = i
			for j >= h && arr[j] < arr[j-h] {
				<-c
				swap(&arr[j], &arr[j-h])
				j -= h
			}
		}
	}
}
