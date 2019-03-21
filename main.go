package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct { //define square type with sideLength field
	sideLength float64 //field
}
type triangle struct { //define triangle type with height and base field
	height float64 //field
	base   float64 //field
}

func main() {
	sq := square{5}       //declare square variable
	tr := triangle{7, 15} //declare triangle variable
	printArea(sq)
	printArea(tr)

}

func (s square) getArea() float64 { //function with receiver of type square
	return s.sideLength * s.sideLength //calculate area of square and return it
}

func (t triangle) getArea() float64 { //function with receiver of type triangle
	return (t.height * t.base) / 2 //calcualte area of square and return it
}

func printArea(sh shape) { //function accept argument type of shape interface
	fmt.Println(sh.getArea()) //print shape's area
}
