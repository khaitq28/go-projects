package maths

func Sum(i int, j int) int {
	return i + j
}
func Sub(i int, j int) int {
	return i - j
}
func Mul(i int, j int) int {
	return i * j
}

func Factorial(i int) int {
	if i == 1 {
		return 1
	}
	return i * Factorial(i-1)
}

func Fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		i, j = j, i+j
		return i
	}
}

func Counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
