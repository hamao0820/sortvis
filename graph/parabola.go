package graph

func parabola() []int {
	var parabola [99]int
	for i := 0; i < 99; i++ {
		parabola[i] = (i - 49) * (i - 49) * 2 / 50
	}
	return parabola[:]
}

func parabolaInverse() []int {
	var parabola [99]int
	for i := 0; i < 99; i++ {
		parabola[i] = 99 - (i-49)*(i-49)*2/50
	}
	return parabola[:]
}
