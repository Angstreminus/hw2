package task2

func task2(a, b, c int) string {
	if (a+b < c) || (a+c < b) || (c+b < a) {
		return "NO"
	}
	return "YES"
}
