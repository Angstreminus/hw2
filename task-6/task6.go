package task6

func task6(arr []int) int {
	dict := make(map[int]int)
	cnt := 0
	for _, elem := range arr {
		_, ok := dict[elem]
		if !ok {
			cnt++
			dict[elem] = elem
		}
	}

	return cnt
}
