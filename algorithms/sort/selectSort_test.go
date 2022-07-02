package sort

import (
	"fmt"
	"testing"
)

func TestSeSort(t *testing.T) {
	res := SeSort([]int{1, 4, 3, 5, 2})
	fmt.Printf("%#v", res)

}
