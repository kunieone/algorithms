package algorithms

// 欧几里得算法，可以算两个数的最大公约数。
func Gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}
