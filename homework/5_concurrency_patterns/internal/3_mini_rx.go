package internal

import "fmt"

type observable struct {
	in  chan int
	out chan int
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
			fmt.Printf("%d just()\n", i+1)
			s.out <- number
		}
	}()

	return observable{in: s.out}
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
			fmt.Printf("%d mapping()\n", i+1)
			i++
			s.out <- m(item)
		}
	}()

	return observable{in: s.out}
}

func (s observable) printer() observable {
	fmt.Println("printer() started")
	var i int
	for item := range s.in {
		fmt.Printf("%d printer()\n", i+1)
		i++
		fmt.Println(item)
	}
	fmt.Println("printer() finished")
	return s
}

func SimpleRxTest() {
	var rx observable
	rx.just(33, 10).mapping(func(i int) int { return i * 3 }).printer()
}
