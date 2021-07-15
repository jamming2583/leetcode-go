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
func reverse1(x int) int {
	var res int
	for x != 0 {
		// 利用go语言的特性，(temp*10)/10 != temp 如果溢出则两边不等
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}
