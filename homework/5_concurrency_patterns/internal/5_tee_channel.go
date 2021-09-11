package internal

import "fmt"

func TeeChannel() {
	//create two channel that can be consumed simultaneously

	tee := func(done <-chan interface{}, data []int) (<-chan int, <-chan int) {
		out1 := make(chan int)
		out2 := make(chan int)

		go func() {
			defer func() {
				close(out1)
				close(out2)
			}()
			for _, item := range data {
				out1Internal := out1
				out2Internal := out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
						return
					case out1Internal <- item:
						out1Internal = nil
					case out2Internal <- item:
						out2Internal = nil
					}
				}
			}
		}()

		return out1, out2
	}

	done := make(chan interface{})
	out1, out2 := tee(done, []int{1, 2, 3, 4})
	for i := range out1 {
		//deadlock if only one channel is read
		fmt.Printf("values: out1=[%d], out2=[%d]\n", i, <-out2)
	}
}
