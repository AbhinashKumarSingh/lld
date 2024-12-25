package lsp

import "fmt"

type Shape interface {
	Area() int64
}

type Rectangle struct {
	Weight, Length int64
}

func (r Rectangle) Area() int64 { return r.Length * r.Weight }

type Square struct {
	Side int64
}

func (s Square) Area() int64 { return s.Side * s.Side }

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}
