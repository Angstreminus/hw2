package task4

func task4(matr [][]int) bool {
	for i := 0; i < len(matr); i++ {
		for j := 0; j < len(matr); j++ {
			if matr[i][j] != matr[j][i] {
				return false
			}
		}
	}
	return true
}
