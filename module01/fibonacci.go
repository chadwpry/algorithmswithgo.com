package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//	Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//
// Examples:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//	Fibonacci(2) => 1
//	Fibonacci(3) => 2
//	Fibonacci(4) => 3
//	Fibonacci(5) => 5
//	Fibonacci(6) => 8
//	Fibonacci(7) => 13
//	Fibonacci(14) => 377

var cache = map[int]int{}

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	fib1, ok := cache[n-1]
	if ok == false {
		cache[n-1] = Fibonacci(n - 1)
		fib1 = cache[n-1]
	}

	fib2, ok := cache[n-2]
	if ok == false {
		cache[n-2] = Fibonacci(n - 2)
		fib2 = cache[n-2]
	}

	return fib1 + fib2
}
