package main

import (
	"algorithms/dataStructures/tree"

	"github.com/gookit/goutil/dump"
)

var (
// p  = fmt.Println
)

func main() {
	arr := []int{2, 5, 6, 6, 7, 4, 9, 1, 3}
	dump.P(tree.New(0, &arr))

}
