package task3

func task3(arr []int) []int {
	res := make([]int, 0)
	size := len(arr) - 1
	res = append(res, arr[size])
	res = append(res, arr[:size]...)
	return res
}
