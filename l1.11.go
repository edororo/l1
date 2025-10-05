package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 8}
	b := []int{4, 5, 6, 7, 1, 8}
	m := make(map[int]struct{})
	peresechenie := make(map[int]struct{})

	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			peresechenie[v] = struct{}{}
		}
	}
	fmt.Println(peresechenie)
}
