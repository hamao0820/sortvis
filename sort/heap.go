package sort

const MAX = 1000

func Heapsort(arr []int) {
	heapify(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = deletemax(arr[:i+1])
	}
}

func heapify(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		downmax(i, arr, len(arr))
	}
}

func downmax(i int, arr []int, n int) {
	j := 2*i + 1

	if j >= n {
		return
	}

	if j+1 < n && arr[j] < arr[j+1] {
		j = j + 1
	}

	if arr[j] > arr[i] {
		swap(&arr[i], &arr[j])
		downmax(j, arr, n)
	}
}

func deletemax(arr []int) int {
	max := arr[0]
	arr[0] = arr[len(arr)-1]
	downmax(0, arr, len(arr)-1)
	return max
}
