package sort

// 计数排序理论上O(n+k),特么太快了  但其实是通过牺牲空间来换取时间的一种算法。并且不能要求数字索引跨度过大。

type Count struct {
	isEmpty bool
	times   int
}

func CSort(arr []int) []int {
	min, max := 0, 0
	for _, e := range arr {
		if e <= min {
			min = e
		}
		if e >= max {
			max = e
		}
	}
	c := makeCounters(max - min + 1)
	for i := 0; i < len(arr); i++ {
		c[arr[i]-min].isEmpty = false
		c[arr[i]-min].times++
	}
	return toSortedArr(c)
}

func makeCounters(len int) []Count {
	counters := make([]Count, len)
	for i := range counters {
		counters[i] = Count{
			isEmpty: true,
			times:   0,
		}
	}
	return counters
}

func toSortedArr(c []Count) []int {
	res := []int{}
	for i, e := range c {
		if !e.isEmpty {
			for ; e.times > 0; e.times-- {
				res = append(res, i)
			}
		}
	}
	return res
}
