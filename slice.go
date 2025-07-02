package main

import (
	"fmt"
	"strings"
)

func names() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)

	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}

func literal() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func withMake() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func printSlice(x []int) {
	fmt.Printf("len:%d cap:%d %v\n", len(x), cap(x), x)
}

func associative() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s,"\n")

	names()
	fmt.Println()

	literal()
	fmt.Println()

	slice := primes[1:4]
	fmt.Println(slice)
	slice = primes[:2]
	fmt.Println(slice)
	slice = primes[1:]
	fmt.Println(slice, "\n")

	printSlice(s)

	var nilSlice []int
	printSlice(nilSlice)
	if nilSlice == nil {
		fmt.Println("nil!\n")
	}

	withMake()
	fmt.Println()

	associative()
	fmt.Println()

	appendSlice()
	fmt.Println()
}