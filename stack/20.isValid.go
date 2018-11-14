package main

import "fmt"

var hashMapReverse = map[string]string{")": "(", "}": "{", "]": "["}
var hashMap = map[string]string{"(": ")", "{": "}", "[": "]"}

/**
 * 1.Traversal string for _, v := range s string(v)
 * 2.define slice := make([]string, 0) not length len(s)
 * 3.slice add element just only slice = append(slice, ele)
 * 4.err: out of indexs
 */
func isValid1(s string) bool {
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

/**
 * 1.generator stack
 * 2.stack => previous next data
 */
type Node struct {
	data     string
	next     *Node
	previous *Node
}

// current pointer
type Stack struct {
	head *Node
}

func NewStack() *Stack {
	return &Stack{head: &Node{}}
}

func (s *Stack) push(data string) {
	node := &Node{data: data, next: nil, previous: s.head}
	s.head.next = node
	s.head = node
}

func (s *Stack) pop() (string, bool) {
	n := s.head
	if s.head == nil || s.head.previous == nil {
		return "", false
	}
	data := n.data
	s.head.previous.next = nil
	s.head = s.head.previous
	return data, true
}

func isValid2(s string) bool {
	var result bool = false
	stack := NewStack()

	for _, v := range s {
		value := string(v)
		if _, ok := hashMap[value]; ok {
			stack.push(value)
		} else if v1, ok := hashMapReverse[value]; ok {
			//if first elem is hashMapReverse then return false
			//judge from hashMapReverse[value] pair
			if stack.head.previous != nil && stack.head.data == v1 {
				stack.pop()
			} else {
				return result
			}
		}
	}
	if stack.head.previous == nil {
		result = true
	}
	return result
}

func main() {
	fmt.Println(isValid2("[[{}]]"))
}
