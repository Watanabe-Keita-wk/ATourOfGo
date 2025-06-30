package main

import "fmt"

func simple() {
	sum := 0
	for i:=0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

// whileもforで記述
func omission() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func loop() {
	for {
	}
}

func main() {
	simple()
	omission()
	// loop()
}