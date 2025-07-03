package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Google"] = Vertex{
		37.42202, -122.08408,
	}
	fmt.Println(m,"\n")

	m2 := make(map[string]int)
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])
	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])
	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])
	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}