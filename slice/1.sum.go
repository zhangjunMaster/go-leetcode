package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	// Arrays have duplicate values，hashMap value is indexs []int
	var hashMap = make(map[int][]int)
	for key, value := range nums {
		key1, ok := hashMap[value]
		if ok {
			hashMap[value] = append(key1, key)
		} else {
			hashMap[value] = []int{key}
		}
	}
	// No matter how many, just take the last one
	for value, key := range hashMap {
		remainInt := target - value
		key1, ok := hashMap[remainInt]
		if ok && key1[0] != key[len(key)-1] {
			result = []int{key[len(key)-1], key1[0]}
			break
		}
	}
	return result
}

func main() {
	nums := []int{3, 2, 1, 1, 4}
	fmt.Println(twoSum(nums, 2))
}
