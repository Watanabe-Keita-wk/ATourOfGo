package main

import "fmt"

func hello() {
	defer fmt.Println("world")

	fmt.Println("Hello")
}

func count() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	hello()

	count()
}