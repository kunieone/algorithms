package tree

import "fmt"

type SegTree struct {
	K string
	V int

	L *SegTree
	R *SegTree
}

func New(start int, arr *[]int) *SegTree {
	s := SegTree{
		K: parse(start, arr),
		V: sum(*arr),
	}
	mid := (len(*arr) + 1) / 2
	l, r := (*arr)[0:(mid)], (*arr)[(mid):]
	if len(*arr) > 1 {
		s.L = New(start, &l)
		s.R = New(start+len(l), &r)
	} else {
		s.L, s.R = nil, nil
	}
	return &s
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

var parse = func(start int, arr *[]int) string {
	return fmt.Sprintf("%v--%v", start, start+len(*arr)-1)
}
