package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	t1Chan := make(chan int)
	t2Chan := make(chan int)

	go Walk(t1, t1Chan)
	go Walk(t2, t2Chan)

	for {
		v1, ok1 := <-t1Chan
		v2, ok2 := <-t2Chan
		switch {
		case !ok1, !ok2:
			return ok1 == ok2
		case v1 != v2:
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}

	test1 := Same(tree.New(1), tree.New(1))
	fmt.Printf("test1:%v\n",test1)
	test2 := Same(tree.New(1), tree.New(2))
	fmt.Printf("test2:%v",test2)
}