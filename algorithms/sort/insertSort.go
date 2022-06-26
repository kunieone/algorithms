package sort

import (
	"fmt"
	"time"
)

func InsertSort(arr []int) []int {
	startTime := time.Now()
	for i := 1; i < len(arr); i++ {
		end := i - 1

		tmp := arr[end+1]
		for end >= 0 && tmp < arr[end] {
			// 现在是tmp是一份原先end+1的拷贝，不用担心数组消失了
			arr[end+1] = arr[end]
			end--
		}
		arr[end+1] = tmp
	}
	elapsedTime := time.Since(startTime) / time.Millisecond // duration in ms
	fmt.Printf("Segment finished in %dms\n", elapsedTime)
	return arr
}

//  [2,3,7,1,9,5,4]
