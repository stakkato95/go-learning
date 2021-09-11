package internal

import (
	"fmt"
	"sync"
)

func ProduceConsumer() {
	producer := func() <-chan int {
		results := make(chan int)

		go func() {
			defer func() {
				close(results)
				fmt.Println("finished production")
			}()

			fmt.Println("start production")
			for i := 0; i < 10; i++ {
				results <- i
			}
		}()

		return results
	}

	consumer := func(data <-chan int) {
		fmt.Println("Prepared for consumption")
		for i := range data {
			fmt.Printf("consumed %d\n", i)
		}
	}

	data := producer()
	consumer(data)
}

func MyProducerConsumerExample() {
	//multiplexing multiple channels to one channel

	producer := func(id int) <-chan interface{} {
		results := make(chan interface{})

		go func() {
			defer func() {
				fmt.Printf(">>>producer %d finished production\n", id)
				close(results)
			}()

			fmt.Printf(">>>producer %d is starting production\n", id)
			for i := 0; i < 10; i++ {
				results <- fmt.Sprintf(">>>producer %d, value %d", id, i)
			}
		}()

		return results
	}

	converter := func(chans ...<-chan interface{}) <-chan interface{} {
		totalData := make(chan interface{})

		var wg sync.WaitGroup

		multiplexer := func(in <-chan interface{}, i int) {
			defer func() {
				fmt.Printf("===finished multiplexing channel #%d\n", i)
				wg.Done()
			}()

			fmt.Printf("===starting multiplexing channel #%d\n", i)
			for item := range in {
				totalData <- item
			}
		}

		wg.Add(len(chans))
		fmt.Println("===starting multiplexing")
		for i, ch := range chans {
			go multiplexer(ch, i+1)
		}
		fmt.Println("===all multiplexing goroutines launched")

		go func() {
			fmt.Println("===waiting for all channels to be consumed")
			wg.Wait()
			fmt.Println("===all channels are consumed")
			close(totalData)
		}()

		return totalData
	}

	consumer := func(data <-chan interface{}) {
		fmt.Println("ready to consume")
		for item := range data {
			fmt.Printf("consumed '%s'\n", item.(string))
		}
		fmt.Println("finished consuming")
	}

	multiplexedData := converter(producer(1), producer(2), producer(3))
	consumer(multiplexedData)
}
