package sort

import (
	"testing"
)

func BenchmarkQSort(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		QSort([]int{1, 2, 5, 7, 4})
	}
	b.StopTimer()
}
