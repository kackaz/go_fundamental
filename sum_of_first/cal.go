package cal

func sumOfFirst(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum
}
