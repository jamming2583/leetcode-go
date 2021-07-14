package main

import "fmt"

func main() {
	nums := []int{4, 7, 2, 15, 9, 3}
	target := 22
	result := twoSum(nums, target)
	fmt.Println(result[0], ",", result[1])
}
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
