package utils

func Squared(num int) int {
	return num * num
}

func Abs[T ~int | ~int64 | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Min[T ~int | ~int64 | ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T ~int | ~int64 | ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Mod(a, b int) int {
	return ((a % b) + b) % b
}
