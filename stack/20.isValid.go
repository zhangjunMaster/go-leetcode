package main

import "fmt"

/**
 * 1.Traversal string for _, v := range s string(v)
 * 2.define slice := make([]string, 0) not length len(s)
 * 3.slice add element just only slice = append(slice, ele)
 * 4.err: out of indexs
 */
func isValid1(s string) bool {
	hashMap := make(map[string]string, 3)
	hashMapReverse := make(map[string]string, 3)
	hashMap["("] = ")"
	hashMap["{"] = "}"
	hashMap["["] = "]"
	hashMapReverse[")"] = "("
	hashMapReverse["}"] = "{"
	hashMapReverse["]"] = "["
	var result bool = false
	slice := make([]string, 0)
	for _, v := range s {
		value := string(v)
		if _, ok := hashMap[value]; ok {
			slice = append(slice, value)
		} else if v1, ok := hashMapReverse[value]; ok {
			if len(slice) != 0 && v1 == slice[len(slice)-1] {
				slice = slice[:len(slice)-1]
			} else {
				return result
			}
		}
	}

	if len(slice) == 0 {
		result = true
	}
	return result
}

func main() {
	fmt.Println(isValid1("]"))
}
