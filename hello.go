package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// }

// func main() {
// 	fmt.Println("Finally")
// 	i := 101
// 	fmt.Println(i)
// }

// func main() {
// 	fmt.Println("1 + 1 = ", 1+1)
// }

// func main() {
// 	fmt.Println(len("Hello World"))
// 	fmt.Println("Hello World"[1])
// 	fmt.Println("Hello " + "World")
// }

// func main() {
// 	fmt.Println(true && true)
// 	fmt.Println(true || false)
// }

// func main() {
// 	var x string = "Hello World"
// 	fmt.Println(x)
// }

// func main() {
// 	x := "Hello World"
// 	fmt.Println(x)
// }

func foo() {
	x := "Hello"
	fmt.Println(x, 12)
}

func foo_2() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func foo_3() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
	}
}

func foo_4() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func print_str(str string) {
	fmt.Println(str)
}

func print_int(number int) {
	fmt.Println(number)
}

func foo_5(x int) {
	switch x {
	case 0: print_str("zero")
	case 1: print_str("one")
	default: print_str("no one")
	}
}

func foo_6() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func foo_7() {
	x := [5]float64{ 98, 93, 77, 82, 83 }
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}

func foo_8() {
	x := [5]float64{ 98, 93, 77, 82, 83 }
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func foo_9() {
	x := [5]float64{ 98, 93, 77, 82, 83 }
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

// func foo_10() {
// 	var x []float64
// 	y := make([]float64, 5, 10)
// }

func foo_11() {
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:4]
	fmt.Println(x)
	y := arr[0:3]
	fmt.Println(y)
}

func foo_12() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func foo_13() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func foo_14() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	a, b := x["key"]
	fmt.Println(a, b)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v:= range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f(n int, ok bool) (int, bool) {
	return n, ok
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func even_generator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	// foo()
	// foo_2()
	// foo_3()
	// foo_4()
	// foo_5(0)
	// foo_5(1)
	// foo_5(7)
	// foo_6()
	// foo_7()
	// foo_8()
	// foo_9()
	// foo_10()
	// foo_11()
	// foo_12()
	// foo_13()
	// foo_14()
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	if n, ok := f(5, false); ok {
		fmt.Println(n, ok)
	} else {
		fmt.Println("no")
	}

	fmt.Println(add(1,2,3,4))

	m := []int{1,2,3}
	fmt.Println(add(m...))

	x := 0
	inc := func() int {
		x++
		return x
	}

	fmt.Println(inc())
	fmt.Println(inc())

	next_even := even_generator()
	for i := 0; i < 3; i++ {
		fmt.Print(next_even())
	}

	
}
