package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, w := range words {
		set[w] = struct{}{}
	}

	fmt.Println("Множество уникальных слов:")
	for w := range set {
		fmt.Println(w)
	}
}
