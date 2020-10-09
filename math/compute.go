package math

// greatest common divisor
func GCD(x int, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// least common multiple
func LCM(x int, y int) int {
	return x * y / GCD(x, y)
}
