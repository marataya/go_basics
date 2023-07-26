package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func (w Wheel) String() string {
	return fmt.Sprintf("Center of wheel: %d, %d; No of spokes: %d", w.Circle.Point.X, w.Y, w.Spokes)
}

func main() {
	w := Wheel{Circle{Point{X: 8, Y: 8}, 5}, 20}
	fmt.Println(w)
}
