package main

import (
	"fmt"

	"unicode"

	"github.com/stakkato95/learning"

	"strconv"
	"strings"
)

type day int

var glob int = -1

//const
const a = 3.14

//enum
type Week int

const (
	Mon Week = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

type person struct {
	name  string
	addr  string
	phone string
}

func main() {
	fmt.Println("hello world")
	fmt.Println(learning.Config())

	//vars
	var monday day = 1
	var x = 100
	y := 100500
	fmt.Printf("day=%v, num=%v, num2=%v, global=%v\n", monday, x, y, glob)

	//pointers
	var xp *int = &x
	fmt.Printf("addr of x = %v, val of x = %v\n", xp, *xp)
	var pointer *int = new(int)
	pointer1 := new(int)
	*pointer = 100501
	*pointer1 = 100502
	fmt.Printf("%v, %v\n", *pointer, *pointer1)

	//type conversion
	var a int32 = 1
	var b int16 = int16(a)
	fmt.Printf("%v, %v\n", a, b)

	//string ops
	var r1 rune = 'a'
	r2 := '1'
	fmt.Printf("is r1 a digit = %v, is r2 a digit = %v\n", unicode.IsDigit(r1), unicode.IsDigit(r2))
	fmt.Println(string(unicode.ToUpper(r1)))
	fmt.Printf("aba > aaa %v\n", strings.Compare("aba", "aaa"))
	fmt.Printf("replacement %v\n", strings.Replace("aaaaaaabcdef", "a", "Z", 2))
	num, err := strconv.Atoi("12345")
	fmt.Printf("num=%v, err=%v\n", num, err)

	//enums
	var mon Week = Mon
	var tue Week = Tue
	var wed Week = 100500
	fmt.Printf("mon=%v, tue=%v, wed=%v\n", mon, tue, wed)

	//control flow
	//for
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	//tagged switch
	var thu Week = Thu
	switch thu {
	case Mon:
		fmt.Println("Monday")
		//no need to write break
	case Tue:
		fmt.Println("Tuesday")
	case Wed:
		fmt.Println("Wednesday")
	case Thu:
		fmt.Println("Thursday")
	case Fri:
		fmt.Println("Friday")
	case Sat:
		fmt.Println("Saturday")
	case Sun:
		fmt.Println("Sunday")
	}
	//tagless switch
	smth := 11
	switch {
	case smth < 11:
		fmt.Println("less")
	case smth > 11:
		fmt.Println("greater")
	default:
		fmt.Println("equal")
	}

	//read user input
	var in string
	_, _ = fmt.Scan(&in)
	fmt.Println(in)

	//array
	ar := [6]int{1, 22, 333, 4444, 55555, 666666}
	printArrayWithSixElements(ar)
	//slice
	slice1 := ar[0:2]
	slice2 := ar[2:6]
	slice3 := []int{6, 7, 8, 9, 10, 11}
	printSlice(slice1)
	printSlice(slice2)
	printSlice(slice3)
	//slice with make
	slice4 := make([]int, 10)
	printSlice(slice4)
	sliceSize := 15
	backingArrayCapacity := 30
	slice5 := make([]int, sliceSize, backingArrayCapacity)
	printSlice(slice5)
	//append
	fmt.Println("append")
	slice6 := make([]int, 2)
	printSlice(slice6)
	slice6 = append(slice6, 4)
	printSlice(slice6)

	//map
	mp1 := make(map[string]string)
	mp2 := map[string]string{
		"one": "eins",
		"two": "zwei",
	}
	delete(mp1, "hi")
	delete(mp2, "three")
	de, p := mp2["three"]
	fmt.Println("\""+de+"\"", p)
	printMap(mp1)
	printMap(mp2)

	//struct
	p1 := person{name: "Alex"}
	fmt.Println(p1.phone)
	fmt.Println(p1)
	p1.addr = "Linz"
	p2 := new(person)
	p2.addr = "Vienna"
	fmt.Println(p2)
}

func printSlice(array []int) {
	for i, v := range array {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%v - %v", i, v)
	}
	fmt.Println()
}

func printArrayWithSixElements(array [6]int) {
	for i, v := range array {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%v - %v", i, v)
	}
	fmt.Println()
}

func printMap(mp map[string]string) {
	for key, val := range mp {
		fmt.Println(key, val)
	}
}
