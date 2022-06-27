package sort

func MSort(arr []int) []int {
	// 合并排序，用的也是分而治之
	if len(arr) <= 1 {
		return arr
	}
	num := len(arr) / 2
	return merge(MSort(arr[:num]), MSort(arr[num:]))
}
func merge(left, right []int) []int {
	var i, j int = 0, 0
	result := []int{}
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
