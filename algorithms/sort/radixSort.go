package sort

import (
	"math"
	"strconv"
)

type Bucket [16][]int

// 获取使用某进制下数组里出现的最大的位数，比如[1,2,3000,234,67] => 4
func getMaxRad(arr []int, base int) int {
	max := 0
	for _, v := range arr {
		l := len(strconv.FormatInt(int64(v), base))
		if max < l {
			max = l
		}
	}
	return max
}

// 把二维bucket展开
func flatBuk(bu Bucket) []int {
	res := []int{}
	for _, v := range bu {
		res = append(res, v...)
	}
	return res
}

// 拿到这个数某进制下某位的数字 0代表各位 如果 256 ，16进制,1 相当于拿到的0xFF 的十位：F 即 15
func GetDigit(num int, dig int, base int) int {
	times := math.Pow(float64(base), float64(dig+1))
	remain := num % int(times)
	return remain / int(math.Pow(float64(base), float64(dig)))
}

func RSort(arr []int, base int) []int {
	maxDig := getMaxRad(arr, base)
	// fmt.Println("maxDig", maxDig)

	return rdx(arr, 0, base, maxDig)
}

func RSort10(arr []int) []int {
	return RSort(arr, 10)
}

func rdx(ar []int, dig int, base int, maxDig int) []int {
	if dig >= maxDig {
		return ar
	}
	var bc Bucket

	for i := 0; i < len(ar); i++ {
		digitValue := GetDigit(ar[i], dig, base)
		bc[digitValue] = append(bc[digitValue], ar[i])
	}
	flat := flatBuk(bc)
	dig++
	return rdx(flat, dig, base, maxDig)
}
