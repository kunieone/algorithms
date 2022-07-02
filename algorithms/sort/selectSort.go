package sort

// 选择排序
func SeSort(arr []int) []int {
	for j := 0; j < len(arr)-1; j++ {
		min := j
		for i := min + 1; i < len(arr); i++ {
			if arr[min] > arr[i] {
				min = i
			}
		}
		arr[j], arr[min] = arr[min], arr[j]

	}
	return arr
}
