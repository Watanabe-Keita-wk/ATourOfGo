package main

import "fmt"

func fibonacci() func() int {
	prevNum := 0
	num := 0
	nextNum := 0
	return func() int {
		if num == 0 && nextNum == 0 {
			nextNum = 1
			return num
		}
		prevNum = num
		num = nextNum
		nextNum = prevNum + num
		return num
	}
}

// ネットで見つけた天才
func genius() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	g := genius()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Println(g())
	}
}