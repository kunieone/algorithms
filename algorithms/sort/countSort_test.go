package sort

import (
	"fmt"
	"testing"
)

func TestCSort(t *testing.T) {
	fmt.Println(CSort([]int{1, 2, 5, 7, 4}))
}
func BenchmarkCSort(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CSort([]int{1, 2, 5, 7, 4})
	}
	b.StopTimer()
}
