package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	//all goroutones finish, when "main()" finished
	go fmt.Println("from goroutine")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("from main")

	//sync pkg
	var wg sync.WaitGroup
	wg.Add(3)
	go job(1, &wg)
	go job(2, &wg)
	go job(3, &wg)
	wg.Wait()
	fmt.Println("main done")

	var c chan int = make(chan int, 10)
	select {
	case c <- 3:
		fmt.Println("written to channel")
	default:
		fmt.Println("error")
	}
	go chanJob(c)
	time.Sleep(100 * time.Millisecond)

	inst := getInstance()
	fmt.Println(inst)
}

func job(n int, wg *sync.WaitGroup) {
	fmt.Println(strconv.Itoa(n) + " finished")
	wg.Done()
}

func chanJob(c chan int) {
	x := <-c
	fmt.Println("received in goroutine " + strconv.Itoa(x))
}

//singleton
var instance string
var once sync.Once

func getInstance() string {
	once.Do(func() {
		instance = "singleton string"
	})

	return instance
}
