package main

import (
	"fmt"
)

func main() {
	s := "abcd"
	for i, v := range s {
		fmt.Println(int(i))

		fmt.Println(string(v))

	}
}
