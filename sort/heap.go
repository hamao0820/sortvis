package sort

const MAX = 1000

func HeapSort(arr []int) {
	heapify(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = deleteMax(arr[:i+1])
	}
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
