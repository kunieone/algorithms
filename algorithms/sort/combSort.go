package sort

func CombSort(data []int) []int {
	l := len(data)

	shrink := 1.3

	gap := l
	sorted := false
	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			sorted = true
			gap = 1
		}
		for i := 0; i < l-gap; i++ {
			if data[i] > data[gap+i] {
				data[i], data[gap+i] = data[gap+i], data[i]
				sorted = false
			}
		}
	}
	return data
}
