package sort

func swap(a, b *int) {
	*a, *b = *b, *a
}

func ValidSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}
