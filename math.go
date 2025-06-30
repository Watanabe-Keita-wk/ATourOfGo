package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値に名前をつけるとreturnだけで返せる。可読性の観点で短い関数でのみ使うことを推奨
func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var j, k int = 1, 2
const Pi = 3.14

var (
	ToBe	bool	   = false
	MaxInt  uint64     = 1<<64 -1
	z		complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println("My faborite number is", rand.Intn(100))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a,b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(j, k, c2, python2, java2)

	l := 3
	fmt.Println(l)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	x,y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Printf("%T %T %T\n", x, y, z)
	fmt.Println(x, y, z)

	i2 := 42
	f2 := 3.142
	g2 := 0.867 + 0.5i
	fmt.Printf("%T %T %T\n", i2, f2, g2)
	fmt.Println(i2, f2, g2)

	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}