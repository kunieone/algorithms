package sort

//快速排序充分利用了分而治之的思路

func QSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// base line
	pivot := arr[0]
	var left, right []int
	for _, e := range arr[1:] {
		if e <= pivot {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}
	return append(QSort(left), append([]int{pivot}, QSort(right)...)...)
}
