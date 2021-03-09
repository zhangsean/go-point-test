package main

import (
	"flag"
	"fmt"
	"time"
)

// Num number struct
type Num struct {
	total int
}

// Func sum method
type Func struct {
	name    string
	handler func(int) int
}

func main() {
	var c = flag.Int("t", 100000, "The target number for sum testing from 1 to")
	flag.Parse()

	fmt.Println("Sum testing from 1 to", *c)
	fmt.Println("")

	fns := []Func{
		{"Normal direct", sumDirect},
		{"Point direct", sumPointDirect},
		{"Normal function", sum},
		{"Point function", sumPoint},
		{"Normal struct", sumStruct},
		{"Point struct", sumPointStruct},
		{"Normal struct function", sumPlusStruct},
		{"Point struct function", sumPlusPointStruct},
	}

	for _, fn := range fns {
		s := time.Now().UnixNano()
		t := fn.handler(*c)
		fmt.Println(fn.name, "sum result:", t)
		e := time.Now().UnixNano()
		ms := float64(e-s) / 1e6
		fmt.Println(fn.name, "sum cost:  ", ms, "ms")
		fmt.Println("")
	}

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

func sumPoint(c int) int {
	var t *int
	a := 0
	t = &a
	for i := 1; i <= c; i++ {
		t = plusPoint(t, &i)
	}
	return *t
}

func plusPoint(a, b *int) *int {
	c := *a + *b
	return &c
}

func sumDirect(c int) int {
	var t int
	for i := 1; i <= c; i++ {
		t += i
	}
	return t
}

func sumPointDirect(c int) int {
	var t *int
	a := 0
	t = &a
	for i := 1; i <= c; i++ {
		*t += i
	}
	return *t
}

func sumStruct(c int) int {
	var t Num
	for i := 1; i <= c; i++ {
		t.total += i
	}
	return t.total
}

func sumPointStruct(c int) int {
	t := &Num{}
	t.total = 0
	for i := 1; i <= c; i++ {
		*&t.total += i
	}
	return t.total
}

func sumPlusStruct(c int) int {
	var t Num
	for i := 1; i <= c; i++ {
		j := Num{}
		j.total = i
		t = plusStruct(t, j)
	}
	return t.total
}

func plusStruct(a, b Num) Num {
	var c Num
	c.total = a.total + b.total
	return c
}

func sumPlusPointStruct(c int) int {
	t := &Num{}
	t.total = 0
	for i := 1; i <= c; i++ {
		j := &Num{}
		j.total = i
		t = plusPointStruct(t, j)
	}
	return t.total
}

func plusPointStruct(a, b *Num) *Num {
	c := &Num{}
	c.total = *&a.total + *&b.total
	return c
}
