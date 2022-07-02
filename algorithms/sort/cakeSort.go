package sort

import (
	"algorithms/algorithms/myutil"
)

// 自创桶排序类 蛋糕排序 切蛋糕 切着切着就好了

func CeSort(arr []int) []int {
	if len(arr) <= 3 {
		return QSort(arr)
	}
	// 1 ---8    8-1 *
	max, min := myutil.GetMaxMin(arr)
	gap := min + ((max - min) / 2)
	small := []int{}
	large := []int{}
	// fmt.Println("midValue", midValue)
	if max == min {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		{
			if (v) <= gap {
				small = append(small, v)
			} else {
				large = append(large, v)

			}
		}
	}

	return append(CeSort(small), CeSort(large)...)
}
