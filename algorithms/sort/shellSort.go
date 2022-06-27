package sort

func ShellSort(arr []int) []int {
	n := len(arr)
	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		for i := step; i < n; i += step {
			inSort(i-step, step, &arr)
		}
	}

	return arr
}

func SSort(list []int) {
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入那么交换元素
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}
func inSort(start int, gap int, arr *[]int) {

	ar := *arr
	for i := start; i >= 0; i -= gap {
		for ar[i+gap] < ar[i] {
			ar[i+gap], ar[i] = ar[i], ar[i+gap]
		}
	}
}
