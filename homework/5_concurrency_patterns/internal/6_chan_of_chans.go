package internal

import "fmt"

func ChanOfChans() {
	toSingleChan := func(done <-chan interface{}, channels <-chan (<-chan int)) <-chan int {
		result := make(chan int)

		go func() {
			defer func() {
				close(result)
			}()

			for channel := range channels {
				for item := range channel {
					select {
					case <-done:
						return
					case result <- item:
					}
				}
			}
		}()

		return result
	}

	createChannels := func() <-chan (<-chan int) {
		allChannels := make(chan (<-chan int))

		go func() {
			defer func() {
				close(allChannels)
			}()

			for i := 0; i < 5; i++ {
				//with normal channel channel
				channel := make(chan int)
				allChannels <- channel
				channel <- i
				close(channel)

				// //with buffered channel channel
				// channel := make(chan int, 1)
				// channel <- i
				// close(channel)
				// allChannels <- channel
			}
		}()

		return allChannels
	}

	allChannels := createChannels()

	done := make(chan interface{})
	result := toSingleChan(done, allChannels)

	for item := range result {
		fmt.Println(item)
	}
}
