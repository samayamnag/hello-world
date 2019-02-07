package main

import(
	"fmt"
	"math"
)


type Geometry interface {
	Area() float64
}


type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}


func (r Rectangle) Area() float64{
	return r.width * r.height
}


func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g Geometry) float64{
	return g.Area()
}

func main()  {
	fmt.Println("Hello world")
	r := Rectangle{12, 2}
	c := Circle{4}

	fmt.Println(r.Area())
	fmt.Println(c.Area())

	fmt.Println(measure(r))

}