package main

import (
	"fmt"
)

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) String() string {
	return fmt.Sprint("Rectangle of width ", r.Width, " and lenght ", r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) String() string {
	return fmt.Sprint("Circle of radius ", c.Radius)
}

type Shape interface {
	String() string
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}
