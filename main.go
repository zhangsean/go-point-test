package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var c = flag.Int("c", 100, "Target number for sum testing from 1 to count")
	flag.Parse()

	fmt.Println("Target number count:", *c)
	fmt.Println("")

	s := time.Now().UnixNano()
	fmt.Println("Normal sum start: ", s)
	t := sum(*c)
	fmt.Println("Normal sum result:", t)
	e := time.Now().UnixNano()
	fmt.Println("Normal sum stop:  ", e)
	ms := float64(e-s) / 1e6
	fmt.Println("Normal sum cost:  ", ms, "ms")
	fmt.Println("")

	s2 := time.Now().UnixNano()
	fmt.Println("Point sum start: ", s2)
	t2 := sumPoint(*c)
	fmt.Println("Point sum result:", *t2)
	e2 := time.Now().UnixNano()
	fmt.Println("Point sum stop:  ", e2)
	ms2 := float64(e2-s2) / 1e6
	fmt.Println("Point sum cost:  ", ms2, "ms")
	fmt.Println("")

}

func sum(c int) int {
	var t int
	for i := 1; i <= c; i++ {
		t = plus(t, i)
	}
	return t
}

func plus(a, b int) int {
	var c int
	c = a + b
	return c
}

func sumPoint(c int) *int {
	var t *int
	a := 0
	t = &a
	for i := 1; i <= c; i++ {
		t = plusPoint(t, &i)
	}
	return t
}

func plusPoint(a, b *int) *int {
	c := *a + *b
	return &c
}
