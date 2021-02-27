package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	iterations  = 7
	goroutines  = 5
	channelSize = 10
)

func main() {
	// //goroutines
	// fmt.Println("hello world")

	// for i := 0; i < goroutines; i++ {
	// 	go doWork(i)
	// }

	// fmt.Scanln()

	// //channels
	// ch1 := make(chan int, channelSize)
	// go func(in <-chan int) {
	// 	// accept only read only channel
	// 	for i := range in {
	// 		fmt.Println(i)
	// 	}
	// }(ch1)

	// ch1 <- 2
	// //close(ch1) // - closing a channel at this point woild cause an error
	// ch1 <- 100500
	// //close(ch1)
	// fmt.Scanln()

	// 	//multiplexing with "select"
	// 	ch2 := make(chan int, 2)
	// 	ch3 := make(chan int, 2)

	// 	//first case will be executed
	// 	ch2 <- 21
	// 	ch2 <- 22
	// 	ch3 <- 31

	// LOOP:
	// 	for {
	// 		select {
	// 		case val := <-ch2:
	// 			fmt.Println("ch2 val =", val)
	// 		// case ch3 <- 1:
	// 		// 	fmt.Println("put val to ch3")
	// 		case val1 := <-ch3:
	// 			fmt.Println("ch3 val =", val1)
	// 		default:
	// 			fmt.Println("default")
	// 			break LOOP
	// 		}
	// 	}

	// //cancelling poducer
	// cancelCh := make(chan struct{})
	// dataChan := make(chan int)

	// go func(cancelCh <-chan struct{}, dataChan chan<- int) {
	// 	val := 0
	// 	for {
	// 		select {
	// 		case <-cancelCh:
	// 			fmt.Println("stopped producer")
	// 			return
	// 		case dataChan <- val:
	// 			val++
	// 		}
	// 	}
	// }(cancelCh, dataChan)

	// for i := range dataChan {
	// 	fmt.Println(i)
	// 	if i == 3 {
	// 		cancelCh <- struct{}{}
	// 		break
	// 	}
	// }

	// //timers
	// timer := time.NewTimer(3 * time.Second)
	// for {
	// 	select {
	// 	case <-timer.C:
	// 		fmt.Println("3 sec timer elapsed")
	// 	case <-time.After(1 * time.Second):
	// 		fmt.Println("1 sec timer elapsed")
	// 	default:
	// 		break
	// 	}
	// }

	// //ticker
	// ticker := time.NewTicker(2 * time.Second)
	// i := 0
	// for tickTime := range ticker.C {
	// 	i++
	// 	fmt.Println(tickTime)
	// 	if i == 3 {
	// 		ticker.Stop()
	// 		break
	// 	}
	// }

	// //call func on time
	// timer := time.AfterFunc(2*time.Second, func() {
	// 	fmt.Println("hello world")
	// })
	// fmt.Scanln()
	// timer.Stop()
	// fmt.Scanln()

	// //multiple parallel requests to services
	// ctx, finish := context.WithCancel(context.Background())
	// out := make(chan int)

	// for i := 0; i < 5; i++ {
	// 	go worker(ctx, i, out)
	// }

	// result := <-out
	// finish()
	// fmt.Println("result delivered by worker #", result)

	// 	//receive results within timeout
	// 	workTime := 50 * time.Millisecond
	// 	ctx, _ := context.WithTimeout(context.Background(), workTime)
	// 	result := make(chan int)
	// 	for i := 0; i < 5; i++ {
	// 		go worker(ctx, i, result)
	// 	}

	// 	totalFound := 0
	// LOOP:
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			break LOOP
	// 		case foundBy := <-result:
	// 			totalFound++
	// 			fmt.Println("result found by", foundBy)
	// 		}
	// 	}
	// 	fmt.Println("total found:", totalFound)

	// 	//async results
	// 	comments := make([]string, 10)
	// 	articles := make([]string, 10)

	// 	commentCh := getComments()
	// 	articleCh := getArticles()

	// LOOP:
	// 	for {
	// 		select {
	// 		case comment := <-commentCh:
	// 			fmt.Printf("#%s\n", comment)
	// 			comments = append(comments, comment)
	// 		case article := <-articleCh:
	// 			fmt.Printf("#%s\n", article)
	// 			articles = append(articles, article)
	// 		case <-time.After(2 * time.Second):
	// 			break LOOP
	// 		}
	// 	}

	// //worker pool
	// workerJobs := make(chan string, 2)
	// workersNum := 4
	// for i := 0; i < workersNum; i++ {
	// 	go startWorker(i, workerJobs)
	// }
	// jobs := []string{"hello1", "world1", "hello2", "world2", "hello3", "world3"}
	// for _, job := range jobs {
	// 	workerJobs <- job
	// }
	// fmt.Scanln()

	// //WaitGroup + goroutines with quotas
	// //quotas limit consumption of resources
	// wg := &sync.WaitGroup{}
	// workersNumber := 10
	// quotaLimit := 2 //only 2 goroutines can execute simultaneously
	// quotaCh := make(chan struct{}, quotaLimit)
	// for i := 0; i < workersNumber; i++ {
	// 	wg.Add(1)
	// 	go startQuotaWorker(i, quotaCh, wg)
	// }
	// wg.Wait()

	//mutex

	//atomic
	for i := 0; i < 1000; i++ {
		go inc()
	}
	fmt.Scanln()
	fmt.Println(totalOperations)
}

func doWork(routineId int) {
	for i := 0; i < iterations; i++ {
		fmt.Printf("%d - %s\n", routineId, strings.Repeat(fmt.Sprint(i), i))
		runtime.Gosched()
	}
}

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Microsecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println(workerNum, "worker done")
		out <- workerNum
	}
}

func getComments() chan string {
	ch := make(chan string, 1)

	go func(ch chan<- string) {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("comment #%d", i)
		}
	}(ch)

	return ch
}

func getArticles() chan string {
	ch := make(chan string, 1)

	go func(ch chan<- string) {
		for i := 0; i < 4; i++ {
			time.Sleep(2 * time.Second)
			ch <- fmt.Sprintf("article #%d", i)
		}
	}(ch)

	return ch
}

func startWorker(workerId int, jobs <-chan string) {
	for job := range jobs {
		fmt.Println("worker #", workerId, "processed '", job, "'")
		runtime.Gosched()
	}
}

func startQuotaWorker(workerId int, quotaCh chan struct{}, wg *sync.WaitGroup) {
	//acquire quota
	quotaCh <- struct{}{}

	fmt.Println("worker #", workerId, " is working")
	time.Sleep(time.Second * 1)

	//release quota
	<-quotaCh
	wg.Done()
}

var totalOperations int32

func inc() {
	////wrong
	//totalOperations++

	//right
	atomic.AddInt32(&totalOperations, 1)
}
