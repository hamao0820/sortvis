package graph

func Ascending() []int {
	var ascending [99]int
	for i := 0; i < 99; i++ {
		ascending[i] = i
	}
	return ascending[:]
}

func Descending() []int {
	var descending [99]int
	for i := 0; i < 99; i++ {
		descending[i] = 99 - i
	}
	return descending[:]
}
