package sort

func ShellSort(arr []int) []int {
	// 数组长度
	n := len(arr)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入那么交换元素
				if arr[j+step] < arr[j] {
					arr[j], arr[j+step] = arr[j+step], arr[j]
					continue
				}
				break
			}
		}
	}
	return arr
}
