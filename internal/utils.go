package internal

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func concat(a, b int) int {
	bTest := b
	bDigits := 0
	for bTest > 0 {
		bDigits++
		bTest /= 10
	}

	return a*int(math.Pow(10, float64(bDigits))) + b
}

func addPair(pair1, pair2 [2]int) [2]int {
	return [2]int{pair1[0] + pair2[0], pair1[1] + pair2[1]}
}

func subPair(pair1, pair2 [2]int) [2]int {
	return [2]int{pair1[0] - pair2[0], pair1[1] - pair2[1]}
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
