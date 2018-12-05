package fib

type FibFunc func() int

func Fibonacci() FibFunc {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
