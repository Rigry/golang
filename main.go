package main

import ("fmt"
		"math/rand"
		"time")

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

type Circle struct {
	x, y, r float64
}

// func circleArea(c *Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// func (c *Circle) area() float64 {
// 	return math.Pi * c.r * c.r
// }

type Rectangle struct {
	x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
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

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// var c Circle
	// c := new(Circle)
	// c.x = 0
	// c.y = 0
	// c.r = 5

	// fmt.Println(c.x, c.y, c.r)
	// fmt.Println(circleArea(c))
	// // fmt.Println(c.area())

	// r := Rectangle{0, 0, 10, 10}
	// fmt.Println(r.area())

	// a := new(Android)
	// a.Name = "Sam"
	// a.Person.Talk()
	// a.Talk()

	// fmt.Println(totalArea(c, &r))
	// x := 5
	// zero(&x)``
	// fmt.Println(x)

	// xPtr := new(int)
	// one(xPtr)
	// fmt.Println(xPtr)
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}