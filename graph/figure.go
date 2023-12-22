package graph

func square() []int {
	var square [99]int
	for i := 0; i < 25; i++ {
		square[i] = 0
	}
	for i := 25; i < 75; i++ {
		square[i] = 75
	}
	for i := 75; i < 99; i++ {
		square[i] = 0
	}
	return square[:]
}

func squareInverse() []int {
	var square [99]int
	for i := 0; i < 25; i++ {
		square[i] = 75
	}
	for i := 25; i < 75; i++ {
		square[i] = 0
	}
	for i := 75; i < 99; i++ {
		square[i] = 75
	}
	return square[:]
}

func step() []int {
	var square [99]int
	for i := 0; i < 25; i++ {
		square[i] = 0
	}
	for i := 25; i < 50; i++ {
		square[i] = 33
	}
	for i := 50; i < 75; i++ {
		square[i] = 66
	}
	for i := 75; i < 99; i++ {
		square[i] = 99
	}
	return square[:]
}

func stepInverse() []int {
	var square [99]int
	for i := 0; i < 25; i++ {
		square[i] = 99
	}
	for i := 25; i < 50; i++ {
		square[i] = 66
	}
	for i := 50; i < 75; i++ {
		square[i] = 33
	}
	for i := 75; i < 99; i++ {
		square[i] = 0
	}
	return square[:]
}

func triangle() []int {
	var square [99]int
	for i := 0; i < 50; i++ {
		square[i] = 2 * i
	}
	for i := 50; i < 99; i++ {
		square[i] = 198 - 2*i - 1
	}
	return square[:]
}

func triangleInverse() []int {
	var square [99]int
	for i := 0; i < 50; i++ {
		square[i] = 198 - 2*(i+49) - 1
	}
	for i := 50; i < 99; i++ {
		square[i] = 2 * (i - 50)
	}
	return square[:]
}
