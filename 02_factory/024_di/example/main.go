package main

import (
	"fmt"

	di "github.com/baqimeilaoban/go-patterb/02_factory/024_di"
)

// A 依赖关系 A -> B -> C
type A struct {
	B *B
}

type B struct {
	C *C
}

type C struct {
	Num int
}

func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

func NewB(c *C) *B {
	return &B{
		C: c,
	}
}

func NewC() *C {
	return &C{
		Num: 1,
	}
}

func main() {
	container := di.New()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}
	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}
}
