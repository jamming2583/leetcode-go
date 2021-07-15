package main

import (
	"math"
)

func main() {
	println(reverse(0))
}

//考的是取余，整除啥的
func reverse(x int) int {
	result := 0
	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			return 0
		}
		tempNum := x % 10
		result = result*10 + tempNum
		x = x / 10
	}
	return result
}

/**
人家写的
*/
//func reverse(x int) (rev int) {
//	for x != 0 {
//		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
//			return 0
//		}
//		digit := x % 10
//		x /= 10
//		rev = rev*10 + digit
//	}
//	return
//}
