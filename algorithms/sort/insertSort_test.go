package sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	c := InsertSort([]int{5, 4, 9, 7, 6, 3, 8, 2, 1})
	if !reflect.DeepEqual(c, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`排序错误！`)
	}
}
func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort([]int{5, 4, 9, 7, 6, 3, 8, 2, 1})
	}
}
