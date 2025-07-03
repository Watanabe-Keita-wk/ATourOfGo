package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	slice := strings.Split(s," ")
	m := make(map[string]int)
	for _, v := range slice {
		_, exist := m[v]
		if exist {
			m[v] += 1
			continue
		}
		m[v] = 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}