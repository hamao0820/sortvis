package graph

import "math"

func sin() []int {
	var sin [99]int
	for i := 0; i < 99; i++ {
		sin[i] = int(math.Floor(math.Sin(math.Pi*float64(i)*2/99)*50 + 50))
	}
	return sin[:]
}

func sin2() []int {
	var sin [99]int
	for i := 0; i < 99; i++ {
		sin[i] = int(math.Floor(math.Sin(math.Pi*float64(i)*4/99)*50 + 50))
	}
	return sin[:]
}

func sin3() []int {
	var sin [99]int
	for i := 0; i < 99; i++ {
		sin[i] = int(math.Floor(math.Sin(math.Pi*float64(i)*6/99)*50 + 50))
	}
	return sin[:]
}

func cos() []int {
	var cos [99]int
	for i := 0; i < 99; i++ {
		cos[i] = int(math.Floor(math.Cos(math.Pi*float64(i)*2/99)*50 + 50))
	}
	return cos[:]
}

func cos2() []int {
	var cos [99]int
	for i := 0; i < 99; i++ {
		cos[i] = int(math.Floor(math.Cos(math.Pi*float64(i)*4/99)*50 + 50))
	}
	return cos[:]
}

func cos3() []int {
	var cos [99]int
	for i := 0; i < 99; i++ {
		cos[i] = int(math.Floor(math.Cos(math.Pi*float64(i)*6/99)*50 + 50))
	}
	return cos[:]
}
