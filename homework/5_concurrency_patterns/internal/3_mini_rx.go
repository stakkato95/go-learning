package internal

import (
	"fmt"
	"time"
)

type observable struct {
	in   chan int
	out  chan int
	done chan interface{}
}

func newObservable() *observable {
	return &observable{done: make(chan interface{})}
}

func (s observable) just(number, times int) observable {
	s.out = make(chan int)

	go func() {
		defer func() {
			fmt.Println("just() finished")
			close(s.out)
		}()

		fmt.Println("just() started")
		for i := 0; i < times; i++ {
			select {
			case <-s.done:
				return
			case s.out <- number:
				fmt.Printf("%d just()\n", i+1)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return observable{in: s.out, done: s.done}
}

func (s observable) mapping(m mapper) observable {
	s.out = make(chan int)

	go func() {
		defer func() {
			fmt.Println("mapping() finished")
			close(s.out)
		}()

		fmt.Println("mapping() started")
		var i int
		for item := range s.in {
			select {
			case <-s.done:
				return
			case s.out <- m(item):
				fmt.Printf("%d mapping()\n", i+1)
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return observable{in: s.out, done: s.done}
}

func (s observable) printer() observable {
	fmt.Println("printer() started")
	var i int
	for item := range s.in {
		select {
		case <-s.done:
			fmt.Println("printer() cancelled")
			return s
		default:
			fmt.Printf("%d printer() - %d\n", i+1, item)
			i++
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("printer() finished")
	return s
}

func SimpleRxTest() {
	rx := newObservable()
	done := rx.done

	go func() {
		fmt.Println("pipeline will be cancelled in 7 sec...")
		time.Sleep(7 * time.Second)
		close(done)
		fmt.Println("pipeline cancelled after 7 sec...")
	}()

	rx.just(33, 10).mapping(func(i int) int { return i * 3 }).printer()
}
