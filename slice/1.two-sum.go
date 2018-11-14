package main

import "fmt"

func twoSum1(nums []int, target int) []int {
	var result []int
outer:
	for key, value := range nums {
		for key1, value1 := range nums {
			if value+value1 == target && key != key1 {
				result = []int{key, key1}
				break outer
			}
		}
	}
	return result
}

func twoSum2(nums []int, target int) []int {
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

func twoSum3(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := hashMap[target-v]; ok {
			return []int{i, j}
		}
		hashMap[v] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 1, 1, 4}
	fmt.Println(twoSum1(nums, 2))
	fmt.Println(twoSum2(nums, 2))
}
