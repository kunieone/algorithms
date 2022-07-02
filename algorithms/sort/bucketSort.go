package sort

import (
	"algorithms/algorithms/myutil"
)

// 桶排序，反正咋都比不过计数排序 因为计数排序其实就是一个桶一个元素。

func BSort(arr []int) []int {
	maxValue, minValue := myutil.GetMaxMin(arr)
	var sz = 5
	type Buck [][]int

	var bk Buck = make(Buck, ((maxValue-minValue)/sz)+1)
	var result []int
	for i := 0; i < len(arr); i++ {
		r := (arr[i] - minValue) / sz
		// fmt.Println(r)
		bk[r] = append(bk[r], arr[i])
	}

	for i := range bk {
		bk[i] = ISort(bk[i])
		result = append(result, bk[i]...)
	}
	return result
}
