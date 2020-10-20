package fibonacci

// Calculate Fibonacci number
func Calculate(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Calculate(n-1) + Calculate(n-2)
}
