package module01

var primes = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
}

func reduce(slices []int, fn func(int, int) int, defaultVal int) int {
	if len(slices) == 0 {
		return defaultVal
	}

	return reduce(slices[1:], fn, fn(defaultVal, slices[0]))
}

func intersect(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}

	var head []int

	for i, val := range b {
		if a[0] == val {
			head = []int{a[0]}
			b = append(b[:i], b[i+1:]...)
			break
		}
	}

	return append(head, intersect(a[1:], b)...)
}

func GCD(a, b int) int {
	aFactors := Factor(primes, a)
	bFactors := Factor(primes, b)

	return reduce(intersect(aFactors, bFactors), func(acc int, i int) int {
		return acc * i
	}, 1)
}
