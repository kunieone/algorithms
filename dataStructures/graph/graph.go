package graph

type Edge struct {
	a      int     //a节点
	b      int     //b节点
	weight float64 //权值
}

func newEdge(a, b int, weight float64) *Edge {
	return &Edge{
		a:      a,
		b:      b,
		weight: weight,
	}
}
func lessWeight(a, b *Edge) bool {
	return a.weight <= b.weight
}

type weightedDenseGraph struct {
	n        int  //节点数
	m        int  //边数
	directed bool //有向图 or 无向图
	graph    [][]*Edge
}
