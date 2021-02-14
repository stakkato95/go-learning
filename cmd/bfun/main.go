package main

import (
	"fmt"
	"strconv"

	"github.com/stakkato95/learning"
)

func main() {
	//args of defer are executed immediately, but call os deferred
	defer fmt.Println("bye!")

	//fun
	fmt.Println("hello")
	sum, str := print(1, 2)
	fmt.Printf("%v %v\n", sum, str)

	//arrays - by val (BAD!!!)
	//ar ref - by ref (BAD!!! constant referenceing and dereferencing)
	//slices - by ref (GOOD!!!)
	ar := [3]int{3, 4, 5}
	arrayHassle(&ar)
	fmt.Println(ar[0])
	slice := ar[:] //ref to ar
	sliceAwesome(slice)
	fmt.Println(ar[0])
	fmt.Println(slice[0])

	//func as var
	var inc func(int) int = increment
	fmt.Println(inc(1))
	inc2 := increment
	fmt.Println(inc2(2))
	fmt.Println(apply(increment, 3))
	//anonymous fun
	fmt.Println(apply(func(x int) int { return x + 1 }, 4))
	fmt.Println(addNumber(2)(4))

	//varargs
	fmt.Println(varargsSum(1, 2, 3, 4, 5, 6))

	//type params
	john := man{name: "John"}
	john.sayName()

	//private & public funs
	fmt.Println(learning.Config())
	p := learning.PublicPerson{PublicName: "pub name"}
	fmt.Println(p)

	//ifaces
	sq := square{side: 3}
	rec := rectangle{sideA: 3, sideB: 4}
	var shp shape = &sq
	fmt.Println(shp.area())
	shp = &rec
	fmt.Println(shp.area())
	//type assertions
	r, ok := shp.(*rectangle)
	if ok {
		fmt.Println("it's a rectangle")
		fmt.Println(r.sideA)
	}
	s, ok := shp.(*square)
	if ok {
		fmt.Println("it's a square")
		fmt.Println(s.side)
	} else {
		fmt.Println("it's not a square")
	}
	//type switch
	switch it := shp.(type) {
	case *rectangle:
		fmt.Println("rectangle" + strconv.Itoa(it.sideA))
	case *square:
		fmt.Println("square " + strconv.Itoa(it.side))
	}
}

func print(x int, y int) (int, string) {
	var str string = strconv.Itoa(x) + "--" + strconv.Itoa(y)
	fmt.Println(x, y)
	return x + y, str
}

func arrayHassle(array *[3]int) {
	dereferenced := (*array)[0]
	(*array)[0] = dereferenced * dereferenced
}

func sliceAwesome(slice []int) {
	slice[0] = slice[0] * slice[0]
}

func increment(n int) int {
	return n + 1
}

func apply(fun func(int) int, arg int) int {
	return fun(arg)
}

func addNumber(n int) func(int) int {
	return func(x int) int { return x + n }
}

func varargsSum(vals ...int) int {
	var sum int = 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

//custom type and fun with receiver type
type man struct {
	name string
}

//pointer receiver is better for big obj
func (this *man) sayName() {
	fmt.Println("My name is " + this.name)
}

//ifaces
type shape interface {
	area() int
	perimeter() int
}

type square struct {
	side int
}

func (this *square) area() int {
	return this.side * this.side
}

func (this *square) perimeter() int {
	return this.side * 4
}

type rectangle struct {
	sideA int
	sideB int
}

func (this *rectangle) area() int {
	return this.sideA * this.sideB
}

func (this *rectangle) perimeter() int {
	return this.sideA*2 + this.sideB*2
}
