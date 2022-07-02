package sort

func CocktailSort(arr []int) []int {
	rightBorder := len(arr) - 1
	leftBorder := 0
	rIndex := 0
	lIndex := 0
	for i := 0; i < len(arr)/2-1; i++ {
		Flag := true
		for j := i; j < rightBorder; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				Flag = false
				rIndex = j
			}
		}
		rightBorder = rIndex
		if Flag {
			break
		}

		for j := rightBorder; j > leftBorder; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				Flag = false
				lIndex = j
			}
		}
		leftBorder = lIndex
		if Flag {
			break
		}
	}
	return arr
}
