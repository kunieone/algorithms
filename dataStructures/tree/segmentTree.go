package tree

type SegTree struct {
	key    string
	value  int
	arrs   []int
	lChild *SegTree
	rChild *SegTree
}

func sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}
