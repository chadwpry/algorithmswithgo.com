package module01

func GCD(a, b int) int {
	var gcd int = 1

	for i := 1; i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}

	return gcd
}
