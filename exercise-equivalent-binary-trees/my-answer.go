package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t != nil {
			walk(t.Left)
			ch <- t.Value
			walk(t.Right)
		}
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			break
		}
		if v2 != v1 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	a := tree.New(1)
	b := tree.New(2)

	go Walk(a, ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(a, a))
	fmt.Println(Same(a, b))
}
