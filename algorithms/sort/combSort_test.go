package sort

import (
	"algorithms/algorithms/myutil"
	"fmt"
	"testing"
)

func TestCombSort(t *testing.T) {
	fmt.Println(CombSort(myutil.GenNoRArr(1000)))
}
