package internal

import (
	"fmt"
	"strings"
	"time"
)

func JoinGoroutine() {
	work := func() <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer close(completed)
			fmt.Println(">>>ready to sleep in work()...")
			time.Sleep(3 * time.Second)
			fmt.Println(">>>wake up in work()...")
		}()

		return completed
	}

	join := work()

	fmt.Println("start waiting in main()...")
	<-join
	fmt.Println("finished waiting in main()...")
}

func WorkCancellationWithDoneChannel() {
	done := make(chan interface{})

	printer := func(done <-chan interface{}, data <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer func() {
				fmt.Println("printer is finished")
				close(completed)
			}()

			for {
				select {
				case s := <-data:
					fmt.Printf("read '%s'\n", s)
				case <-done:
					fmt.Println("'case <-done' triggered")
					return
				}
			}
		}()

		return completed
	}

	producer := func() <-chan string {
		data := make(chan string)

		go func() {
			defer func() {
				fmt.Println("finished production")
				close(data)
			}()

			for i := 1; i < 10; i++ {
				data <- strings.Repeat(fmt.Sprint(i), i)
				time.Sleep(1 * time.Second)
			}
		}()

		return data
	}

	join := printer(done, producer())

	go func() {
		fmt.Println("printer will be stopped in 3 sec...")
		time.Sleep(3 * time.Second)
		close(done)
		fmt.Println("printer is stopped after 3 sec...")
	}()

	fmt.Println("waiting in main()...")
	<-join
	fmt.Println("finished waiting in main()...")
}
