package main

import (
  "fmt"
  "math"
)

func main() {
  rect := Rectangle{20, 40}
  circ := Circle{10}

  fmt.Println("Area of Rectangle is", rect.area())
  fmt.Println("Area of Circle is", circ.area())
}


type Shape interface {
  area() float64
}

type Rectangle struct {
  width float64
  height float64
}

type Circle struct {
  radius float64
}

func (r Rectangle) area() float64{
  return r.width * r.height
}

func (r Circle) area() float64{
  return math.Pi * math.Pow(r.radius, 2)
}

func getArea(shape Shape) float64 {
  return shape.area()
}



