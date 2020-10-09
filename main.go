package main

import ("fmt"; "math")

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}
func (r *Rectangle) area_r() float64 {
	fmt.Println("Unknown")
	return 1
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}



func main() {
	// var c Circle
	c := new(Circle)
	c.x = 0
	c.y = 0
	c.r = 5

	fmt.Println(c.x, c.y, c.r)
	fmt.Println(circleArea(c))
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area_r())

	a := new(Android)
	a.Name = "Sam"
	a.Person.Talk()
	a.Talk()
	// x := 5
	// zero(&x)
	// fmt.Println(x)

	// xPtr := new(int)
	// one(xPtr)
	// fmt.Println(xPtr)
}