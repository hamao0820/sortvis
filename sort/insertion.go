package sort

func InsertionSort(arr []int) {
	var i, j int

	for i = 1; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j] < arr[j-1] {
			swap(&arr[j], &arr[j-1])
			j--
		}
	}
}

func InsertionSortAsync(arr []int, c chan struct{}) {
	var i, j int

	for i = 1; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j] < arr[j-1] {
			<-c
			swap(&arr[j], &arr[j-1])
			j--
		}
	}
}
