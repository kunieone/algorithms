package myutil

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GenNoRArr(length int) []int {
	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr = append(arr, i)
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}

func Rival(length int, cb func([]int) []int) string {
	testArr := GenNoRArr(length)
	sT := time.Now()
	cb(testArr)
	eT := time.Since(sT)
	t := float64(eT.Microseconds())
	fmt.Printf("%4f", (t*math.Log2(float64(length)))/float64(length))
	fmt.Println(" ")
	return fmt.Sprintf("%.4f", (t*math.Log2(float64(length)))/float64(length))
}

func RivalTime(length int, cb func([]int) []int) time.Duration {
	testArr := GenNoRArr(length)
	sT := time.Now()
	cb(testArr)
	eT := time.Since(sT)
	return eT
}

func GetMaxMin(arr []int) (max int, min int) {
	max, min = arr[0], arr[0]
	for _, v := range arr {

		if v >= max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}
