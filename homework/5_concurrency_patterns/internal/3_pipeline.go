package internal

import "fmt"

type mapper func(int) int

func Pipeline() {
	repeated := repeat(111, 5)
	mapped := mapping(func(i int) int { return i * 3 }, repeated)
	stringed := toString(mapped)
	doPrint(stringed)
}

func repeat(number, times int) <-chan int {
	result := make(chan int)

	go func() {
		defer func() {
			fmt.Println("repeat() finished")
			close(result)
		}()

		fmt.Println("repeat() starting")
		for i := 0; i < times; i++ {
			result <- number
		}
	}()

	return result
}

func mapping(m mapper, in <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		defer func() {
			fmt.Println("mapping() finished")
			close(result)
		}()

		fmt.Println("mapping() starting")
		for item := range in {
			result <- m(item)
		}
	}()

	return result
}

func toString(in <-chan int) <-chan string {
	result := make(chan string)

	go func() {
		defer func() {
			fmt.Println("toString() finished")
			close(result)
		}()

		fmt.Println("toString() starting")
		for item := range in {
			result <- fmt.Sprint(item)
		}
	}()

	return result
}

func doPrint(in <-chan string) {
	fmt.Println("doPrint() starting")
	for item := range in {
		fmt.Println(item)
	}
	fmt.Println("doPrint() finished")
}
