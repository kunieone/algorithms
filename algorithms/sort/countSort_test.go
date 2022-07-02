package sort

import (
	"algorithms/algorithms/myutil"
	"fmt"
	"testing"
)

func TestCSort(t *testing.T) {
	cS := myutil.RivalTime(300_00_000, CSort)
	fmt.Println("cS", cS)
}
func BenchmarkCSort(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cS := myutil.RivalTime(300_000, CSort)
		fmt.Println("cS", cS)

	}
	b.StopTimer()
}
