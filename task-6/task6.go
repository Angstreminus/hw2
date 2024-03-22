package task6

func task6(arr []int) int {
	dict := make(map[int]struct{})
	cnt := 0
	for _, elem := range arr {
		//! интересный момент, ревьюер спасибо, сам бы не сразу додумался
		if _, ok := dict[elem]; ok {
			continue
		}
		cnt++
		dict[elem] = struct{}{}
	}
	return cnt
}
